package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wrap(handler func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := handler(c)

		if err != nil {
			var httpErr *HTTPError

			if errors.As(err, &httpErr) {
				c.JSON(httpErr.Code, gin.H{"error": httpErr.Error()})
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// Check if the response is a SuccessResponse type
		if successResp, ok := response.(SuccessResponse); ok {
			c.JSON(successResp.Status, successResp.Data)
		} else {
			c.JSON(http.StatusOK, response) // Default to 200 OK
		}
	}
}
