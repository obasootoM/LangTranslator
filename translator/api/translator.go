package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/translator/db/sqlc"
)

// translator signup
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

type LoginTranslatorResponse struct {
	SessionID   uuid.UUID                  `form:"session_id" json:"session_id"`
	AccessToken string                     `form:"accesstoken" json:"accesstoken"`
	Translator  TranslatorRegisterResponse `form:"translator" json:"translator"`
	RefreshToken          string           `form:"refresh_token" json:"refresh_token"`
	AccessTokenExpiredAt  time.Time        `form:"access_token_expired_at"`
	RefreshtokenExpiredAt time.Time        `form:"refresh_token_expired_at" json:"refresh_token_expired_at"`

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

// translator login
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
	accessstoken, accessPayload, err := server.token.CreateToken(req.Email, server.config.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	refreshToken, refreshPayload, err := server.token.CreateToken(req.Email, server.config.RefreshDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	session, err := server.store.CreateSession(ctx,db.CreateSessionParams{
        ID: refreshPayload.ID,
		Email: trans.Email,
		RefreshToken: refreshToken,
		UserAgent: "",
		IsBlocked: false,
		ExpiresAt: refreshPayload.ExpiredAt,

	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	arg := LoginTranslatorResponse{
		SessionID: session.ID,
		Translator:  NewTranslator(trans),
		AccessToken: accessstoken,
		RefreshtokenExpiredAt: accessPayload.ExpiredAt,
		RefreshToken: refreshToken,
		AccessTokenExpiredAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, arg)
}

// get translator by their email
type GetEmailrequest struct {
	Translator TranslatorRegisterResponse `form:"translator" json:"translator"`
}
type GetEmail struct {
	Email string `form:"email" json:"email"`
}

func (server *Server) getTranslator(ctx *gin.Context) {
	var req GetEmail
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	getemail, err := server.store.GetTranslator(ctx, req.Email)
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := GetEmailrequest{
		Translator: NewTranslator(getemail),
	}
	ctx.JSON(http.StatusOK, arg)
}

// delete translator
type DeleteTrans struct {
	Email string `form:"email" json:"email" `
}

func (server *Server) delete(ctx *gin.Context) {
	var req DeleteTrans

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	err := server.store.DeleteTranslator(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully delete",
	})

}


func (server *Server) logout(ctx *gin.Context) {
	var req LoginTranslatorResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	session, err := server.store.GetSession(ctx, req.SessionID)
	if err !=nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	err = server.store.UpdateSession(ctx, db.UpdateSessionParams{
		ID: session.ID,
		IsBlocked: true,
	})
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully logged out",
	})
	
}
