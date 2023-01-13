package api

import (
	"database/sql"
	_ "log"
	"net/http"
	_ "net/smtp"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

type RegisterClient struct {
	Firstname  string `form:"firstname" json:"firstname" xml:"firstname"  binding:"required,alphanum"`
	Secondname string `form:"secondname" json:"secondname" xml:"secondname"  binding:"required,alphanum"`
	Email      string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}
type RegisterResponse struct {
	FirstName  string    `form:"firstname" json:"firstname"`
	SecondName string    `form:"firstname" json:"secondname"`
	Email      string    `form:"email" json:"email"`
	CreateAt   time.Time `from:"createat" json:"createat"`
	UpdatedAt  time.Time `form:"updatedat" json:"updatedat"`
}

func NewClientResponse(client db.Client) RegisterResponse {
	return RegisterResponse{
		FirstName:  client.FirstName,
		SecondName: client.SecondName,
		Email:      client.Email,
		UpdatedAt:  client.UpdatedAt,
		CreateAt:   client.CreatedAt,
	}
}

func (server *Server) createClient(ctx *gin.Context) {

	var req RegisterClient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashpassword, err := config.Hashpassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateClientParams{
		FirstName:  req.Firstname,
		SecondName: req.Secondname,
		Email:      req.Email,
		Password:   hashpassword,
	}
	clients, err := server.store.CreateClient(ctx, arg)
	if err != nil {
		if pkErr, ok := err.(*pq.Error); ok {
			switch pkErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	resp := NewClientResponse(clients)

	ctx.JSON(http.StatusOK, resp)
}

type LoginClientRequest struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}

type LoginClientResponse struct {
	SessionID             uuid.UUID        `form:"session_id" json:"session_id"`
	AccessToken           string           `form:"accesstoken" json:"accesstoken"`
	Client                RegisterResponse `form:"client"`
	RefreshToken          string           `form:"refresh_token" json:"refresh_token"`
	AccessTokenExpiredAt  time.Time        `form:"access_token_expired_at"`
	RefreshtokenExpiredAt time.Time        `form:"refresh_token_expired_at" json:"refresh_token_expired_at"`
}

func NewClient(client db.Client) RegisterResponse {
	clients := RegisterResponse{
		FirstName:  client.FirstName,
		SecondName: client.SecondName,
		Email:      client.Email,
		UpdatedAt:  client.UpdatedAt,
		CreateAt:   client.CreatedAt,
	}
	return clients
}

func (server *Server) loginClient(ctx *gin.Context) {
	var req LoginClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	clients, err := server.store.GetClient(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = config.CheckPassword(req.Password, clients.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accesstoken, accessPayload, err := server.token.CreateToken(req.Email, server.config.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	refreshToken, refreshPayload, err := server.token.CreateToken(req.Email, server.config.RefreshDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Email:        clients.Email,
		RefreshToken: refreshToken,
		UserAgent:    "",
		IsBlocked:    false,
		ClientIp:     "",
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := LoginClientResponse{
		SessionID:             session.ID,
		Client:                NewClient(clients),
		AccessToken:           accesstoken,
		RefreshToken:          refreshToken,
		AccessTokenExpiredAt:  accessPayload.ExpiredAt,
		RefreshtokenExpiredAt: refreshPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, arg)
}

type GetClientEmail struct {
	Email string `form:"email" json:"email"`
}

func (server *Server) getClientEmail(ctx *gin.Context) {
	var req GetClientEmail

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	getclientE, err := server.store.GetClient(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, getclientE)
}

type DeleteClient struct {
	Email string `form:"email" json:"email"`
}

func (server *Server) deleteclient(ctx *gin.Context) {
	var req DeleteClient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	err := server.store.DeleteClient(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})

}

func (server *Server) logout(ctx *gin.Context) {
	var req LoginClientResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	get, err := server.store.GetSession(ctx, req.SessionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = server.store.UpdateSession(ctx, db.UpdateSessionParams{
		ID:        get.ID,
		IsBlocked: true,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"mesage": "successfully logout",
	})
}

type ChangePassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (server *Server) changePassword(ctx *gin.Context) {
	var req ChangePassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	hashpassword, err := config.Hashpassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	get, err := server.store.GetClient(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.UpdateClientParams{
		Email:    get.Email,
		Password: hashpassword,
	}
	err = server.store.UpdateClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "password updated",
	})
}
