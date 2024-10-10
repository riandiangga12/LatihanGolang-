package handler

import (
	"LatihanGolang/aggregator/service"
	"LatihanGolang/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IAggregatorHandler is an interface that defines the handler methods for the aggregator service
type IAggregatorHandler interface {
	CreateUserAndWallet(c *gin.Context)
	GetUserAndWallet(c *gin.Context)
	TopUpWallet(c *gin.Context)
	TransferWallet(c *gin.Context)
	GetTransactions(c *gin.Context)
}

type aggregatorHandler struct {
	aggregatorService service.IAggregatorService
}

// NewAggregatorHandler creates a new instance of AggregatorHandler
func NewAggregatorHandler(aggregatorService service.IAggregatorService) IAggregatorHandler {
	return &aggregatorHandler{
		aggregatorService: aggregatorService,
	}
}

// CreateUserAndWallet is a handler method to create user and wallet
func (h *aggregatorHandler) CreateUserAndWallet(c *gin.Context) {
	var user struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User and Wallet
	entityUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	userResponse, walletResponse, err := h.aggregatorService.CreateUserAndWallet(c, &entityUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":   userResponse,
		"wallet": walletResponse,
	})

}

// GetUserAndWallet is a handler method to get user and wallet by user ID
func (h *aggregatorHandler) GetUserAndWallet(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, wallet, err := h.aggregatorService.GetUserAndWallet(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"wallet": wallet,
	})
}

// TopUpWallet is a handler method to top up wallet balance
func (h *aggregatorHandler) TopUpWallet(c *gin.Context) {
	var topUp struct {
		ID     int     `json:"id" binding:"required"`
		Amount float32 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&topUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet, err := h.aggregatorService.TopUpWallet(c, topUp.ID, topUp.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wallet": wallet.Message,
	})
}

// TransferWallet is a handler method to transfer wallet balance
func (h *aggregatorHandler) TransferWallet(c *gin.Context) {
	var transfer struct {
		From   int     `json:"from" binding:"required"`
		To     int     `json:"to" binding:"required"`
		Amount float32 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet, err := h.aggregatorService.TransferWallet(c, transfer.From, transfer.To, transfer.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wallet": wallet.Message,
	})
}

// GetTransactions is a handler method to get transactions by user ID
func (h *aggregatorHandler) GetTransactions(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactions, err := h.aggregatorService.GetTransactions(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}
