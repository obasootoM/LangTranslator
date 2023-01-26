package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type Pages struct {
	SourceLanguage    string `json:"sourcelanguage" form:"sourcelanguage" binding:"required"`
	TargetLanguage    string `json:"targetlanguage" form:"targetlanguage" binding:"required"`
	Profession        string `json:"profession" form:"profession" binding:"required,oneof=standard professional premium"`
	Category          string `json:"category" form:"category" binding:"required"`
	Field             string `json:"field" form:"field" binding:"required"`
	Duration          string `json:"duration" form:"duration" binding:"required"`
	AdditionalService string `json:"additionalservice" form:"additionalservice" binding:"required"`
}

func (server *Server) createPage(ctx *gin.Context) {
	var req Pages
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	arg := db.CreateClientPagesParams{
		SourceLanguage:    req.SourceLanguage,
		TargetLanguage:    req.TargetLanguage,
		Profession:        req.Profession,
		Category:          req.Category,
		Field:             req.Field,
		Duration:          req.Duration,
		AdditionalService: req.AdditionalService,
	}

	page, err := server.store.CreateClientPages(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, page)
}

type File struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required"`
}

func (server *Server) createFile(ctx *gin.Context) {
	var req File
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	err := godotenv.Load()
	if err != nil{
        log.Fatal("Error loading .env file")
	}
	file, header, _ := ctx.Request.FormFile("file")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	uploadFile,openErr := header.Open()
    if openErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
            "message":"cannot open s3",
		})
		return
	}
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("file-image-bucket"),
		Key:    aws.String(header.Filename),
		Body:   uploadFile,
		ACL: "public-read",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	buf := bytes.NewBuffer(nil)
	numbers, err := io.Copy(buf, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot read file",
		})
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("number of words :%v,%v",numbers, result.Location))
}
