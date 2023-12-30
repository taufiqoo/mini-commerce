package handler

import (
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/transaction"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetAllTransactionHandler(c *gin.Context) {
	userID, exists := c.Get("currentUser")
	if !exists {
		responseErr := helper.APIFailure(404, "not found", "user ID not found in context")
		c.JSON(404, responseErr)
		return
	}
	userIDInt, ok := userID.(int)
	if !ok {
		responseErr := helper.APIFailure(400, "bad request", "invalid user ID type in context")
		c.JSON(400, responseErr)
		return
	}

	transactions, err := h.service.GetAllTransaction(userIDInt)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get all transaction", transactions)
	c.JSON(200, response)
}

func (h *transactionHandler) GetTransactionDetailHandler(c *gin.Context) {
	transactionID, _ := strconv.Atoi(c.Param("transactionId"))

	transaction, err := h.service.GetTransactionDetail(transactionID)
	if err != nil {
		responseErr := helper.APIFailure(404, "Not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get transaction detail", transaction)
	c.JSON(200, response)
}

func (h *transactionHandler) CreateTransactionHandler(c *gin.Context) {
	userID, exists := c.Get("currentUser")
	if !exists {
		responseErr := helper.APIFailure(404, "not found", "user ID not found in context")
		c.JSON(404, responseErr)
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		responseErr := helper.APIFailure(400, "bad request", "invalid user ID type in context")
		c.JSON(400, responseErr)
		return
	}

	var input entity.TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIFailure(400, "All field are required", gin.H{"message": splitErr})
		c.JSON(400, responseErr)
		return
	}

	newTransaction, err := h.service.SaveNewTransaction(userIDInt, input)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse(201, "success create new transaction", newTransaction)
	c.JSON(201, response)
}

// BUAT ADMIN
func (h *transactionHandler) GetAllUserTransactionHandler(c *gin.Context) {
	transaction, err := h.service.GetAllUserTransaction()
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get all user transaction", transaction)
	c.JSON(200, response)
}

func (h *transactionHandler) DeleteTransactionHandler(c *gin.Context) {
	transactionID, _ := strconv.Atoi(c.Param("transactionId"))

	transaction, err := h.service.DeleteTransaction(transactionID)
	if err != nil {
		responseErr := helper.APIFailure(404, "transaction not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success delete transaction", transaction)
	c.JSON(200, response)
}
