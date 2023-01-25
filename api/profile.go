package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type Profile struct {
	Name           string `json:"name" form:"name" binding:"required"`
	Gender         string `json:"gender" form:"gender" binding:"required,oneof=female male"`
	Email          string `json:"email" form:"email" binding:"required"`
	PhoneNumber    string `json:"phonenumber" form:"phonenumber" binding:"required"`
	Country        string `json:"country" form:"country" binding:"required"`
	NativeLanguage string `json:"nativelanguage" form:"nativelanguage" binding:"required"`
	AddressLine    string `jso:"addressline" form:"addressline" binding:"required"`
}

func (server *Server) createProfile(ctx *gin.Context) {
	var req Profile
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateProfileParams{
		Name:           req.Name,
		Gender:         req.Gender,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		AddressLine:    req.AddressLine,
		Country:        req.Country,
		NativeLanguage: req.NativeLanguage,
	}

	profile, err := server.store.CreateProfile(ctx, arg)
	if err != nil {
		if pkErr, ok := err.(*pq.Error); ok {
			switch pkErr.Code.Name() {
			case "UNIQUE_VIOLATION":
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, profile)
}

type ListProfile struct {
	PageId   int32 `form:"page_id" json:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" json:"page_size" binding:"required,min=5,max=10"`
}
type ProfileResponse struct {
	Name        string    `form:"name"`
	Email       string    `form:"email"`
	PhoneNumber string    `form:"phone_number"`
	Time        time.Time `form:"time"`
}

//	type ProfileRequest struct {
//		profile  ProfileResponse
//	}
func NewProfile(profile db.Profile) ProfileResponse {
	profiles := ProfileResponse{
		Name:        profile.Name,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
		Time:        profile.CreatedAt,
	}
	return profiles
}

func (server *Server) listProfile(ctx *gin.Context) {
	var req ListProfile
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	arg := db.ListProfileParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	list, err := server.store.ListProfile(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// profiles := ProfileRequest{
	//     profile: NewProfile(list[]),
	// }

	ctx.JSON(http.StatusOK, list)
}

type Images struct {
	Image  multipart.FileHeader `form:"image" json:"image"`
}

func (server *Server) createImage(ctx *gin.Context) {
	var req Images
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	image, _ := ctx.FormFile("image")
	fmt.Println(image.Filename)
	err := ctx.SaveUploadedFile(image,"upload/images/ " + image.Filename)
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":"image upload successful",
	})
}
