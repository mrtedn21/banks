package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDatabase()

	handler := NewHandler()
	router := gin.Default()

	router.POST("/account", handler.CreateAccount)
	router.GET("/account/:id", handler.GetAccount)
	router.PUT("/account/:id", handler.UpdateAccount)
	router.DELETE("/account/:id", handler.DeleteAccount)

	router.Run()
}
