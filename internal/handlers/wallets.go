package handlers

import (
	"fmt"
	"github.com/farhodm/ewallet/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (h *Handler) CheckExistingWallet(ctx *gin.Context) {
	var wallet models.Wallet
	err := h.DB.Where("user_id =?", ctx.Param("userID")).First(&wallet).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "wallet not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "wallet exists"})
}

type Replenish struct {
	UserID uuid.UUID `json:"id" binding:"required"`
	Amount uint64    `json:"amount" binding:"required"`
}

func (h *Handler) ReplenishWallet(ctx *gin.Context) {
	var data Replenish
	if err := ctx.ShouldBindJSON(&data); err != nil {
		log.Println("cannot parse data:", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	var wallet models.Wallet
	err := h.DB.Where("user_id =?", data.UserID).First(&wallet).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "wallet not found"})
		return
	}

	if wallet.Balance+data.Amount > getMaxBalance(wallet.Type) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "maximum balance exceeded",
		})
		return
	}

	wallet.Balance += data.Amount
	transaction := models.Transaction{
		WalletID:  wallet.ID,
		Amount:    data.Amount,
		CreatedAt: time.Now(),
	}
	if err = h.DB.Save(&wallet).Error; err != nil {
		log.Println("cannot save wallet to DB:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if err = h.DB.Create(&transaction).Error; err != nil {
		log.Println("cannot create transaction:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "replenished successfully"})
}

func (h *Handler) GetTransactions(ctx *gin.Context) {
	var wallet models.Wallet
	err := h.DB.Where("user_id =?", ctx.Param("userID")).Preload("Transactions").First(&wallet).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "wallet not found"})
		return
	}

	now := time.Now()
	year, month, _ := now.Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	var totalTransactions int64
	var totalAmounts uint64

	err = h.DB.Model(&models.Transaction{}).
		Where("wallet_id =? and created_at >=? and created_at <=?", wallet.ID, firstOfMonth, time.Now()).Count(&totalTransactions).
		Select("sum(amount)").Row().Scan(&totalAmounts)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"totalAmounts":      fmt.Sprintf("%.2f", totalAmounts),
		"totalTransactions": totalTransactions,
	})
}

func (h *Handler) GetBalanceWallet(ctx *gin.Context) {
	var wallet models.Wallet
	err := h.DB.Where("user_id =?", ctx.Param("userID")).First(&wallet).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "wallet not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": fmt.Sprintf("%.2f", wallet.Balance)})
}

func getMaxBalance(accountType string) uint64 {
	if accountType == "identified" {
		return 100_000_000
	}
	return 10_000_00
}
