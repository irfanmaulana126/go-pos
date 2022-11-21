package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/irfanmaulana126/go-pos/handlers"
	"github.com/irfanmaulana126/go-pos/middlewares"
	"github.com/irfanmaulana126/go-pos/repositorys"
	"github.com/irfanmaulana126/go-pos/services"
)

func NewRouteCustomer(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositoryCustomer(db)
	service := services.NewServiceCustomer(repository)
	handler := handlers.NewHandlerCustomer(service)

	route := router.Group("/api/v1/customer")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "outlet": true}))

	router.GET("/api/v1/customer/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
