package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store   // store is a field of type *db.Store
	router *gin.Engine // router is a field of type *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// TODO: add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	// we should not query and return all of them in a single API call
	router.GET("/accounts", server.listAccounts)
	server.router = router

	return server
}

// errorResponse is a helper function to generate error response.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
