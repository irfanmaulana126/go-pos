package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/irfanmaulana126/go-pos/handlers"
	"github.com/irfanmaulana126/go-pos/middlewares"
	"github.com/irfanmaulana126/go-pos/repositorys"
	"github.com/irfanmaulana126/go-pos/services"
)

func NewRouteProduct(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositoryProduct(db)
	service := services.NewServiceProduct(repository)
	handler := handlers.NewHandlerProduct(service)

	route := router.Group("/api/v1/product")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "merchant": true}))

	router.GET("/api/v1/product/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
	route.GET("/outlet/:id", handler.HandlerResultByOutlet)
}
