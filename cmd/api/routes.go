package main

import (
	"github.com/farhodm/alif-test/internal/handlers"
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

	//router.Use(middleware.Authenticate())
	wallet := router.Group("/wallet")
	{
		wallet.GET("/check", h.CheckExistingWallet)
		wallet.GET("/replenish", h.ReplenishWallet)
		wallet.GET("/transactions", h.GetTransactions)
		wallet.GET("/balance", h.GetBalanceWallet)
	}
	return router
}
