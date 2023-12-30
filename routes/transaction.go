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
	r.GET("/transactions/:transactionId", middleware.Authentication(authService), middleware.Authorization("user"), transactionHandler.GetTransactionDetailHandler)
	r.POST("/transactions", middleware.Authentication(authService), middleware.Authorization("user"), transactionHandler.CreateTransactionHandler)

	r.GET("/transactions-user", middleware.Authentication(authService), middleware.Authorization("admin"), transactionHandler.GetAllUserTransactionHandler)
	r.DELETE("/transactions/:transactionId", middleware.Authentication(authService), middleware.Authorization("admin"), transactionHandler.DeleteTransactionHandler)
}
