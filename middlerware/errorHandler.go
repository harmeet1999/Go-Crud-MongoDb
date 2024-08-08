package errorHandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if an error occurred
		err := c.Errors.Last()
		if err != nil {
			// Log the error
			log.Printf("Error occurred: %v", err.Error())

			// Prepare custom HTTP response JSON
			httpResponse := gin.H{
				"error": "Internal Server Error",
			}

			// Abort the request with a 500 status code and return JSON response
			c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
			return
		}
	}
}
