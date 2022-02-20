package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vukieuhaihoa/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new HTTP server and setting routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	// add routes for router
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
