package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) ValidateBody(c *gin.Context, data any) error {
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	if err := validate.Struct(&data); err != nil {
		return err
	}

	return nil
}

func (app *Config) WriteJSON(c *gin.Context, status int, data any) {
	c.JSON(status, JsonResponse{Status: status, Message: "Success", Data: data})
}

func (app *Config) ErrorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	c.JSON(statusCode, JsonResponse{Status: statusCode, Message: err.Error()})
}
