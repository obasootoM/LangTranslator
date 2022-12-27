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

type ClientRequest struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=7"`
}

type LoginClientRequest struct {
	Client RegisterResponse `form:"client"`
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
	err = config.CheckPassword(req.Password, client.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := LoginClientRequest{
		Client: NewClient(client),
	}
	ctx.JSON(http.StatusOK, arg)
}
