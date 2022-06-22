package httputil

import "github.com/gin-gonic/gin"

func NewError(c *gin.Context, code int, message string) {
	err := HttpError{
		Code:    code,
		Message: message,
	}
	c.JSON(code, err)

}

type HttpError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
