package error

import "net/http"

// CommonError struct
type CommonError struct {
	Code         int         `json:"code"`
	ResponseCode int         `json:"responseCode,omitempty"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// BadRequest struct
type BadRequest struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewBadRequest
func NewBadRequest() BadRequest {
	errObj := BadRequest{}
	errObj.Message = "Bad Request"
	errObj.Code = http.StatusBadRequest

	return errObj
}

// NotFound struct
type NotFound struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewNotFound() NotFound {
	errObj := NotFound{}
	errObj.Message = "NotFound"
	errObj.Code = http.StatusNotFound

	return errObj
}

// Unauthorized struct
type Unauthorized struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewUnauthorized() Unauthorized {
	errObj := Unauthorized{}
	errObj.Message = "Unauthorized"
	errObj.Code = http.StatusUnauthorized

	return errObj
}

// Conflict struct
type Conflict struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewConflict() Conflict {
	errObj := Conflict{}
	errObj.Message = "Conflict"
	errObj.Code = http.StatusConflict

	return errObj
}

// InternalServerError struct
type InternalServerError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewInternalServerError() InternalServerError {
	errObj := InternalServerError{}
	errObj.Message = "Internal Server Error"
	errObj.Code = http.StatusInternalServerError

	return errObj
}
