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
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			return
		}

		if res, ok := response.(*APIResponse); ok {
			c.JSON(res.Status, res.Data)
		} else {
			c.JSON(http.StatusOK, response)
		}
	}
}
