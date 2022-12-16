package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/obasootom/langtranslator/db/sqlc"
)





type Server struct {
     store *db.Store
	 router *gin.Engine
}



func NewServer(store *db.Store) *Server {
	server := Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/signup",server.createClient)
	router.POST("/loginin",server.loginClient)
	server.router = router
	return &server
}

func errorResponse(err error) gin.H {
  return gin.H{"err":err.Error()}
}
func (server Server) StartTls(address string) error {
	return server.router.RunTLS(address,"cert.perm","key.perm")
}