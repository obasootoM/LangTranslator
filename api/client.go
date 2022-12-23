package api

import (
	"database/sql"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type RegisterClient struct {
	Firstname   string `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,alphanum"`
	Secondname  string `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,alphanum"`
	Email       string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" xml:"phonenumber" binding:"required,min=11"`
	Currency    string `form:"currency" json:"currency" xml:"currency" binding:"required,oneof= USD GBP EUR CHF JPY TRY"`
	Language    string `form:"language" json:"language" xml:"language" binding:"required,oneof= EN FR ES DE NL TR AR ZH"`
	Password    string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
	Time        string `form:"time" json:"time" binding:"required"`
}

func (server *Server) createClient(ctx *gin.Context) {

	var req RegisterClient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashpassword, err := config.Hashpassword(req.Password) 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	arg := db.CreateClientParams{
		FirstName:   req.Firstname,
		SecondName:  req.Secondname,
		Email:       req.Email,
		PhoneNumber: string(req.PhoneNumber),
		Language:    req.Language,
		Password:    hashpassword,
		Time:        req.Time,
		Currency:    req.Currency,
	}
	clent, err := server.store.CreateClient(ctx, arg)
	if err != nil {
		if pkErr, ok := err.(*pq.Error); ok {
			switch pkErr.Code.Name(){
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, clent)
}

type ClientRequest struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}

type ClientResponse struct {
	FirstName   string    `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,firstname"`
	SecondName  string    `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,secondname"`
	Email       string    `form:"email" json:"email" xml:"email"  binding:"required,email"`
	PhoneNumber string    `form:"phonenumber" xml:"phonenumber" binding:"required,min=11"`
	Language    string    `form:"language" json:"language" xml:"language" binding:"required,language"`
	Password    string    `form:"password" json:"password" xml:"password" binding:"required,min=7"`
	Time        string    `json:"time" form:"time" binding:"required"`
	UpdatedAt   time.Time `json:"updatedat"`
	CreateAt    time.Time `json:"createdat"`
}

type LoginClientRequest struct {
	Client ClientResponse `form:"client"`
}

func NewClient(client db.Client) ClientResponse {
	clients := ClientResponse{
		FirstName:   client.FirstName,
		SecondName:  client.SecondName,
		Email:       client.Email,
		Language:    client.Language,
		PhoneNumber: client.PhoneNumber,
		Password:    client.Password,
		UpdatedAt:   client.UpdatedAt,
		CreateAt:    client.CreatedAt,
		Time:        client.Time,
	}
	return clients
}

func (server *Server) loginClient(ctx *gin.Context) {
	var req ClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	client, err := server.store.GetClient(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := LoginClientRequest{
		Client: NewClient(client),
	}
	ctx.JSON(http.StatusOK, arg)
}
