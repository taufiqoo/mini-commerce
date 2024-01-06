package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mini-commerce/config"
	"mini-commerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// db := config.Connection()

	// db.AutoMigrate(&migration.User{})
	// db.AutoMigrate(&migration.Product{})
	// db.AutoMigrate(&migration.Cart{})
	// db.AutoMigrate(&migration.Address{})
	// db.AutoMigrate(&migration.Transaction{})

	// db.Migrator().DropTable(&migration.User{})
	// db.Migrator().DropTable(&migration.Product{})
	// db.Migrator().DropTable(&migration.Cart{})
	// db.Migrator().DropTable(&migration.Address{})
	// db.Migrator().DropTable(&migration.Transaction{})

	routes.UserRoute(r)
	routes.ProductRoute(r)
	routes.CartRoute(r)
	routes.AddressRoute(r)
	routes.TransactionRoute(r)

	go consumeUpdatedTransaction()

	r.Run()
}

func consumeUpdatedTransaction() {
	conn, err := config.RabbitMQConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	ctx := context.Background()
	updateTransaction, err := ch.ConsumeWithContext(
		ctx,
		"update-transaction",
		"update-transaction",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err)
	}

	for message := range updateTransaction {
		var transactionData map[string]interface{}
		if err := json.Unmarshal(message.Body, &transactionData); err != nil {
			fmt.Println("Failed to parse message body:", err)
			continue
		}

		id, idExists := transactionData["id"].(float64)
		statusPayment, statusExists := transactionData["status"].(string)
		if !idExists || !statusExists {
			fmt.Println("Failed to retrieve ID or Status")
			continue
		}

		fmt.Printf("ID: %.0f, Status: %s\n", id, statusPayment)
	}

}
