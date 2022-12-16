package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type Register struct {
	FirstName  string `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,firstname"`
	SecondName string `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,secondname"`
	Email      string `form:"email" json:"email" xml:"email"  binding:"required, email"`
	Language   string `form:"language" json:"language" xml:"language" binding:"required,language"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}

func (server *Server) CreateClient(ctx *gin.Context) {

	var req Register
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateClientParams{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Email:      req.Email,
		Language:   req.Language,
		Password:   req.Password,
	}
	clent, err := server.store.CreateClient(ctx, arg)
	if err != nil {
		if pkErr, ok := err.(*pq.Error); ok {
			switch pkErr.Code {
			case "unique violation":
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
	Email    string `form:"email" json:"email" xml:"email"  binding:"required, email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}

type ClientResponse struct {
	FirstName  string `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,firstname"`
	SecondName string `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,secondname"`
	Email      string `form:"email" json:"email" xml:"email"  binding:"required, email"`
	Language   string `form:"language" json:"language" xml:"language" binding:"required,language"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
	PasswordChangedAt time.Time `form:"passwordchangeat" json:"passwordchangeat"`
	UpdatedAt time.Time `json:"updatedat"`
	ChangedAt time.Time `json:"changedat"`
}
