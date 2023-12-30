package middleware

import (
	"log"
	"mini-commerce/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(authService Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": "unauthorized user"})

			c.AbortWithStatusJSON(401, responseError)
			return
		}

		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": err.Error()})

			c.AbortWithStatusJSON(401, responseError)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		print(claims)

		if !ok || !token.Valid {
			responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": "unauthorized user"})

			c.AbortWithStatusJSON(401, responseError)
			return
		}

		ID := int(claims["user_id"].(float64))
		Role := string(claims["role"].(string))

		c.Set("currentUser", ID)
		log.Println(ID, " << ID")
		c.Set("role", Role)
		log.Println(Role, " << Role")
	}
}
