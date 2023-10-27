package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ali423/hexagonal/internal/util"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type response struct {
	context *gin.Context
	rs      gin.H
}

func Response(c *gin.Context) *response {
	return &response{
		context: c,
		rs:      gin.H{},
	}
}

func (r response) Message(msg string) response {
	r.rs["message"] = msg
	return r
}

func (r response) Payload(payload interface{}) response {
	r.rs["data"] = payload
	return r
}

func (r response) Status(statusCode int) {
	r.context.JSON(statusCode, r.rs)
}

func ResponseSuccess(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, gin.H{
		"data": payload,
	})
}

func ResponseFail(c *gin.Context, statusCode int, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := translateErrors(validationErrors)
		c.JSON(statusCode, gin.H{
			"errors": errors,
		})
		return
	}

	value := reflect.ValueOf(err)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		count := value.Len()
		errors := make([]ValidationError, 0)
		for i := 0; i < count; i++ {
			validationErrors := value.Index(i).Interface().(validator.ValidationErrors)
			translatedErrors := translateErrors(validationErrors)
			errors = append(errors, translatedErrors...)
		}
		c.JSON(statusCode, gin.H{
			"errors": errors,
		})
		return
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		c.JSON(statusCode, gin.H{
			"message": util.INVALID_INPUT_DATA_FORMAT,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": err.Error(),
	})
}

func translateErrors(validationErrors validator.ValidationErrors) []ValidationError {
	var errors []ValidationError
	for _, fieldErr := range validationErrors {
		message := ""
		switch fieldErr.Tag() {
		case "required":
			message = fmt.Sprintf(util.VALIDATION_ERROR_REQUIRED, translateFieldName(fieldErr.Field()))
		case "gte":
			message = fmt.Sprintf(util.VALIDATION_ERROR_GREATER_THAN, translateFieldName(fieldErr.Field()), translateFieldName(fieldErr.Param()))
		case "required_without":
			message = fmt.Sprintf(util.VALIDATION_ERROR_REQUIRED_WITHOUT, translateFieldName(fieldErr.Field()), translateFieldName(fieldErr.Param()))
		default:
			message = fmt.Sprintf(util.VALIDATION_ERROR_DEFAULT, translateFieldName(fieldErr.Field()))
		}
		errors = append(errors, ValidationError{
			Field: fieldErr.Field(),
			Error: message,
		})
	}
	return errors
}

func translateFieldName(fieldName string) string {
	if val, ok := util.FIELDS[fieldName]; ok {
		return val
	} else {
		return fieldName
	}
}

func stringToUint(parameter string) (*uint, error) {
	if param, err := strconv.Atoi(parameter); err != nil {
		return nil, err
	} else {
		uintParam := uint(param)
		return &uintParam, nil
	}
}

func stringToInt(parameter string) (*int, error) {
	if param, err := strconv.Atoi(parameter); err != nil {
		return nil, err
	} else {
		return &param, nil
	}
}
