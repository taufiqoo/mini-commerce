package middleware

import (
	"log"
	"mini-commerce/helper"

	"github.com/gin-gonic/gin"
)

func Authorization(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": "role not found in token"})
			c.AbortWithStatusJSON(401, responseError)
			return
		}

		userRole, ok := role.(string)
		if !ok {
			responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": "invalid role type"})
			c.AbortWithStatusJSON(401, responseError)
			return
		}

		allowed := false
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			responseError := helper.APIFailure(403, "Forbidden", gin.H{"message": "user role not authorized"})
			c.AbortWithStatusJSON(403, responseError)
			return
		}

		log.Println(userRole, " << Authorized Role")
		c.Next()
	}
}
