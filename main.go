package main

import (
	"Bookstore/controller"
	"Bookstore/funcii"
	"github.com/gin-gonic/gin"
)

var (
	bookShelf   funcii.Bookshelf      = funcii.New()
	bookControl controller.Controller = controller.New(bookShelf)
)

func main() {
	server := gin.Default()

	server.GET("/GetAll", func(ctx *gin.Context) {
		ctx.JSON(200, bookControl.GetAll())
	})

	server.GET("/GetSortedAcs", func(ctx *gin.Context) {
		ctx.JSON(200, bookControl.GetSortedAscending())
	})

	server.GET("/GetSortedDes", func(ctx *gin.Context) {
		ctx.JSON(200, bookControl.GetSortedDescending())
	})
	server.POST("/Add", func(ctx *gin.Context) {
		ctx.JSON(200, bookControl.Add(ctx))
	})
	server.Run(":8080")
}
