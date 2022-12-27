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

type RegisterTranslator struct {
	FirstName  string `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,alphanum"`
	SecondName string `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,alphanum"`
	Email      string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}
type TranslatorRegisterResponse struct {
	FirstName  string    `form:"firstname" json:"firstname"`
	SecondName string    `form:"firstname" json:"secondname"`
	Email      string    `form:"email" json:"email"`
	CreateAt   time.Time `from:"createat" json:"createat"`
	UpdatedAt  time.Time `form:"updatedat" json:"updatedat"`
}

func NewTranslatorResponse(trans db.Translator) TranslatorRegisterResponse {
	return TranslatorRegisterResponse{
		FirstName:  trans.FirstName,
		SecondName: trans.SecondName,
		Email:      trans.Email,
		CreateAt:   trans.CreatedAt,
		UpdatedAt:  trans.UpdatedAt,
	}
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
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Email:      req.Email,
		Password:   hashpassword,
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
	resp := NewTranslatorResponse(tranlator)
	ctx.JSON(http.StatusOK, resp)
}

type TranslatorRequest struct {
	Email    string `form:"email" xml:"email" json:"email" binding:"required,email"`
	Password string `form:"password" xml:"password" json:"password" binding:"required,min=7"`
}

type LoginTranslatorRequest struct {
	Translator TranslatorRegisterResponse `form:"translator" json:"translator"`
}

func NewTranslator(trans db.Translator) TranslatorRegisterResponse {
	translators := TranslatorRegisterResponse{

		FirstName:  trans.FirstName,
		SecondName: trans.SecondName,
		Email:      trans.Email,
		CreateAt:   trans.CreatedAt,
		UpdatedAt:  trans.UpdatedAt,
	}

	return translators
}

func (server *Server) loginTranslator(ctx *gin.Context) {
	var req TranslatorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	trans, err := server.store.GetTranslator(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = config.CheckPassword(req.Password, trans.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := LoginTranslatorRequest{
		Translator: NewTranslator(trans),
	}
	ctx.JSON(http.StatusOK, arg)
}
