package routes

import (
	"fmt"
	"historico_consumo/App/Controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	e := gin.Default()
	configureRoutes(e)
	err := e.Run(":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on port 8080")
}

func configureRoutes(e *gin.Engine) {
	g := e.Group("/api")
	routes(g)
}

func routes(g *gin.RouterGroup) {
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	nota_fiscal := Controllers.NewNotaFiscalController()
	g.GET("/nota-fiscal", nota_fiscal.Index)
	g.POST("/nota-fiscal", nota_fiscal.Create)
	//g.POST("/nota-fiscal", controllers.CreateNotaFiscal)
}
