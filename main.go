package main

import (
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

	r.Run()
}
