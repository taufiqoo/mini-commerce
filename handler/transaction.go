package handler

import (
	"context"
	"encoding/json"
	"mini-commerce/config"
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/transaction"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"

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
		responseErr := helper.APIFailure(404, "not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(201, "success create new transaction", newTransaction)
	c.JSON(201, response)
}

func (h *transactionHandler) PaymentTransactionHandler(c *gin.Context) {
	transactionID, _ := strconv.Atoi(c.Param("transactionId"))

	var paymentInput entity.PaymentInput
	if err := c.ShouldBindJSON(&paymentInput); err != nil {
		responseErr := helper.APIFailure(400, "Bad request", "Invalid request body")
		c.JSON(400, responseErr)
		return
	}

	transactionDetails, err := h.service.GetTransactionDetail(transactionID)
	if err != nil {
		responseErr := helper.APIFailure(404, "Not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	if paymentInput.Nominal != float64(transactionDetails.TotalPrice) {
		responseErr := helper.APIFailure(400, "Bad request", "Invalid nominal total price")
		c.JSON(400, responseErr)
		return
	}

	var updatedTransaction entity.Transaction
	updatedTransaction, err = h.service.UpdateStatusPaidTranscation(transactionID)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	err = h.publishUpdatedTransaction(updatedTransaction)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse(200, "success payment transaction", nil)
	c.JSON(200, response)
}

func (h *transactionHandler) publishUpdatedTransaction(updatedTransaction entity.Transaction) error {
	conn, err := config.RabbitMQConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	ctx := context.Background()
	updatedTransactionJSON, _ := json.Marshal(updatedTransaction)
	err = ch.PublishWithContext(ctx, "transaction", "update-transaction", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(updatedTransactionJSON),
	})
	if err != nil {
		return err
	}

	return nil
}

// ONLY ADMIN CAN ACCESS THIS HANDLER
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
