package httpapi

import (
	inventory "practice-golang/inventory/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *inventory.Store
	router *gin.Engine
}

func NewServer(store *inventory.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/products", server.checkOrderHandler)
	router.GET("/products/:sku", server.getProductHandler)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
