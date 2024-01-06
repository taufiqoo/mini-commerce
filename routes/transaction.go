package routes

import (
	"mini-commerce/handler"
	"mini-commerce/layer/transaction"
	"mini-commerce/middleware"

	"github.com/gin-gonic/gin"
)

var (
	transactionRepo    = transaction.NewRepository(DB)
	transactionService = transaction.NewService(transactionRepo)
	transactionHandler = handler.NewTransactionHandler(transactionService)
)

func TransactionRoute(r *gin.Engine) {
	r.GET("/transactions", middleware.Authentication(authService), middleware.Authorization("user"), transactionHandler.GetAllTransactionHandler)
	r.GET("/transaction/:transactionId", middleware.Authentication(authService), middleware.Authorization("user"), transactionHandler.GetTransactionDetailHandler)
	r.POST("/transaction", middleware.Authentication(authService), middleware.Authorization("user"), transactionHandler.CreateTransactionHandler)
	r.POST("/transaction/payment/:transactionId", middleware.Authentication(authService), middleware.Authorization("user"), transactionHandler.PaymentTransactionHandler)

	r.GET("/transactions-user", middleware.Authentication(authService), middleware.Authorization("admin"), transactionHandler.GetAllUserTransactionHandler)
	r.DELETE("/transaction/:transactionId", middleware.Authentication(authService), middleware.Authorization("admin"), transactionHandler.DeleteTransactionHandler)
}
