package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type RegisterTranslator struct {
	FirstName  string `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,alphanum"`
	SecondName string `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,alphanum"`
	Email      string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}

func (server *Server) createTranslator(ctx *gin.Context) {
	var req RegisterTranslator
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashpassword, err := config.Hashpassword(req.Password)
	if err != nil {
       ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	   return
	}
	arg := db.CreateTranslatorParams{
		FirstName:          req.FirstName,
		SecondName:         req.SecondName,
		Email:              req.Email,
		Password:           hashpassword,
	}

	tranlator, err := server.store.CreateTranslator(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusMethodNotAllowed, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tranlator)
}
