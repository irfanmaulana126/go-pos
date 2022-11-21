package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/irfanmaulana126/go-pos/handlers"
	"github.com/irfanmaulana126/go-pos/middlewares"
	"github.com/irfanmaulana126/go-pos/repositorys"
	"github.com/irfanmaulana126/go-pos/services"
)

func NewRouteTransaction(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositoryTransaction(db)
	service := services.NewServiceTransaction(repository)
	handler := handlers.NewHandlerTransaction(service)

	route := router.Group("/api/v1/transaction")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(map[string]bool{"admin": true, "merchant": true, "outlet": true}))

	router.GET("/api/v1/transaction/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update:id", handler.HandlerUpdate)
}
