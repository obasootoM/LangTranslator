package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

}
