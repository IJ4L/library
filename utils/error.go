package utils

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"library.com/data/response"
)

func HandleError(ctx *gin.Context, statusCode int, statusMessage string, err error) {
	ctx.JSON(statusCode, response.ErorrResponse{
		Code:    statusCode,
		Status:  statusMessage,
		Message: err.Error(),
	})
	ctx.Abort()
}

func ValidateFunc(obj interface{}, validate *validator.Validate) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Validate:", r)
		}
	}()

	if err := validate.Struct(obj); err != nil {
		errorValid := err.(validator.ValidationErrors)
		var errs error

		for _, e := range errorValid {
			snp := e.StructNamespace()
			errmgs := errorTagFunc(obj, snp, e.Field(), e.ActualTag())
			if errmgs != nil {
				errs = errors.Join(errs, fmt.Errorf("%w", errmgs))
			} else {
				errs = errors.Join(errs, fmt.Errorf("%w", e))
			}
		}

		if errs != nil {
			return errs
		}
	}

	return nil
}

func errorTagFunc(obj interface{}, namespace, field string, tag string) error {
	switch tag {
	case "required":
		return fmt.Errorf("[%s] field is a required or invalid", field)
	default:
		return nil
	}
}
