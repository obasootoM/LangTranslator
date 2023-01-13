package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/db/sqlc"
	"github.com/obasootom/langtranslator/token"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
	token  token.Maker
	config config.Config
}

func NewServer(store *db.Store, config config.Config) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token %d", err)
	}
	server := Server{
		config: config,
		store:  store,
		token:  token,
	}
	router := gin.Default()
	router.POST("/client/signup", server.createClient)
	router.POST("/client/login", server.loginClient)
	router.POST("/client/profile", server.createProfile)
	router.POST("/client/chanPassword", server.changePassword)
	auth := router.Group("/admin")
	auth.GET("/client/get", server.getClientEmail)
	auth.DELETE("/client/delete", server.deleteclient)
	router.GET("/client/logout", server.logout)
	auth.GET("/client/list", server.listProfile)
	auth.POST("/client/order", server.createOrder)

	server.router = router
	return &server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"err": err.Error()}
}
func (server Server) StartTls(address string) error {
	return server.router.RunTLS(address, "cert.pem", "key.pem")
}
func (server Server) Start(address string) error {
	return server.router.Run(address)
}
