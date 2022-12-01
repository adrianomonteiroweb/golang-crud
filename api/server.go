package api

import (
	db "github.com/adrianomonteiroweb/golang-crud/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.StoreExec
	router *gin.Engine
}

func InstanceServer(store *db.StoreExec) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/product", server.createProduct)
	router.PUT("/product", server.updateProduct)
	router.DELETE("/product/:id", server.deleteProduct)
	router.GET("/product/:id", server.getProduct)
	router.GET("/products", server.getProducts)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"API has one error:": err.Error()}
}