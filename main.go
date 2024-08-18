package main

import (
	"github.com/gin-gonic/gin"
)

func main (){
	// Iniciando o Router utilizando as configurações Default do gin
	router *gin.Engine := gin.Default()
	// Definindo uma Rota
	router.GET(relativePath: "/ping", handlers ... : func(c *gin.Context) {
		c.JSON(code: 200, obj: gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a nossa api
	router.Run(addr ... : ":8080") 
}