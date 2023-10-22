package main

import (
	"github.com/farhodm/ewallet/internal/handlers"
	"github.com/farhodm/ewallet/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func routes(db *gorm.DB) *gin.Engine {
	h := handlers.NewHandler(db)

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	})

	router.Use(middleware.AuthMiddleware(db))
	wallet := router.Group("/wallet")
	{
		wallet.GET("/:id/check", h.CheckExistingWallet)
		wallet.POST("/replenish", h.ReplenishWallet)
		wallet.GET("/:id/transactions", h.GetTransactions)
		wallet.GET("/:id/balance", h.GetBalanceWallet)
	}
	return router
}
