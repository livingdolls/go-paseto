package response

import (
	"errors"
	"net/http"

	"github.com/livingdolls/go-paseto/internal/core/entity"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var errorStatusMap = map[error]int{
	entity.ErrInternal:                   http.StatusInternalServerError,
	entity.ErrDataNotFound:               http.StatusNotFound,
	entity.ErrConflictingData:            http.StatusConflict,
	entity.ErrInvalidCredentials:         http.StatusUnauthorized,
	entity.ErrUnauthorized:               http.StatusUnauthorized,
	entity.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	entity.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	entity.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	entity.ErrInvalidToken:               http.StatusUnauthorized,
	entity.ErrExpiredToken:               http.StatusUnauthorized,
	entity.ErrForbidden:                  http.StatusForbidden,
	entity.ErrNoUpdatedData:              http.StatusBadRequest,
	entity.ErrInsufficientStock:          http.StatusBadRequest,
	entity.ErrInsufficientPayment:        http.StatusBadRequest,
	entity.ErrNoMatchPassword:            http.StatusUnauthorized,
}

func HandleErrorResponse(c *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]

	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRespon := newErrorResponse(errMsg)
	c.AbortWithStatusJSON(statusCode, errRespon)

}

func RequestValidationError(c *gin.Context, err error) {
	errMsg := parseError(err)
	errRespons := newErrorResponse(errMsg)
	c.JSON(http.StatusBadRequest, errRespons)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// response represents a response body format
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func HandleSuccessResponse(c *gin.Context, data any) {
	resp := newResponse(true, "Success", data)

	c.JSON(http.StatusOK, resp)
}

func HandleSuccessResponseCreated(c *gin.Context, data any) {
	resp := newResponse(true, "Success", data)

	c.JSON(http.StatusCreated, resp)
}
