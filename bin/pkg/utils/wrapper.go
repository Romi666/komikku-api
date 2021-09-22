package utils

import (
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	httpError "komiku-srapper/bin/pkg/http-error"
	"net/http"
)

// Result common output
type Result struct {
	Data  interface{}
	Error interface{}
}


//BaseWrapperModel data structure
type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}

//Response function
func Response(data interface{}, message string, code int, c echo.Context) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	}
	LogSuccess(data, message)
	return c.JSON(code, result)
}

//ResponseError function
func ResponseError(err interface{}, c echo.Context) error {
	errObj := getErrorStatusCode(err)
	result := BaseWrapperModel{
		Success: false,
		Data:    errObj.Data,
		Message: errObj.Message,
		Code:    errObj.Code,
	}
	LogError(errObj.Data, errObj.Message)
	return c.JSON(errObj.ResponseCode, result)
}

func getErrorStatusCode(err interface{}) httpError.CommonError {
	errData := httpError.CommonError{}

	switch obj := err.(type) {
	case httpError.BadRequest:
		errData.ResponseCode = http.StatusBadRequest
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Unauthorized:
		errData.ResponseCode = http.StatusUnauthorized
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Conflict:
		errData.ResponseCode = http.StatusConflict
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.NotFound:
		errData.ResponseCode = http.StatusNotFound
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.InternalServerError:
		errData.ResponseCode = http.StatusInternalServerError
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	default:
		errData.Code = http.StatusConflict
		return errData
	}
}

// LogSuccess is a function of success process that only generate log to console
func LogSuccess(detail interface{}, msg string) {
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
	}).Info(msg)
}

// LogError is a function of failed process that only generate log to console
func LogError(detail interface{}, msg string) {
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
	}).Error(msg)
}

// LogActivity is a function to create format log contains method, url, and request body. It's created after service hit.
func LogActivity(method interface{}, url interface{}, req interface{}, msg string) {
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"method":  method,
		"url":     url,
		"request": req,
	}).Info(msg)
}

