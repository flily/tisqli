package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Details []string    `json:"details,omitempty"`
	Data    interface{} `json:"data"`
}

func DefaultHandler404(c *gin.Context) {
	response := &Response{
		Code:    404,
		Success: false,
		Message: fmt.Sprintf("route to '%s %s' not found",
			c.Request.Method, c.Request.RequestURI),
	}

	c.JSON(404, response)
}

func DefaultHandlerFatal(c *gin.Context) {
	defer func(c *gin.Context) {
		response := &Response{
			Code:    500,
			Success: false,
			Message: "internal server error",
			Details: []string{
				"critical server error occurred",
			},
		}

		if info := recover(); info != nil {
			response.Details = append(response.Details, fmt.Sprintf("%v", info))
			c.JSON(500, response)
		}

	}(c)

	c.Next()
}

func Success(data interface{}) (*Response, error) {
	response := &Response{
		Code:    200,
		Success: true,
		Data:    data,
	}

	return response, nil
}

func NotFound(key interface{}) (*Response, error) {
	response := &Response{
		Code:    404,
		Success: true,
		Message: fmt.Sprintf("not found, key = %v", key),
	}

	return response, nil
}

func errorResponse(c *gin.Context, code int, message string, details ...string) {
	response := &Response{
		Code:    code,
		Success: false,
		Message: message,
		Details: details,
	}

	c.JSON(code, response)
}
func MakeHandler(handler func(c *gin.Context) (*Response, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := handler(c)
		if err != nil {
			log.Printf("Error: %v", err)
			errorResponse(c, 400, "bad request", err.Error())
			return
		}

		c.JSON(response.Code, response)
	}
}
