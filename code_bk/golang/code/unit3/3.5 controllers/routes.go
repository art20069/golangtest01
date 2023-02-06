package main

import (
	"codezard-pos/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)
}
