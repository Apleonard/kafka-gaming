package api

import (
	"kafka-gaming/api/handlers"

	"github.com/gin-gonic/gin"
)

func Router(listenAddr string) error {

	router := gin.New()
	router.POST("/kocak", handlers.NewHandler().Test)

	return router.Run(":8080")
}
