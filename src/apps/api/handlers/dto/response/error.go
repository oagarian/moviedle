package response

import (
	"net/http"
	"moviedle/src/core/domain/errors"
	"github.com/labstack/echo/v4"
)


type ErrorMessage struct {
	Code          int            `json:"status_code,omitempty"`
	Message       string         `json:"message"`
	InvalidFields []InvalidField `json:"invalid_fields,omitempty"`
	isInternal    bool
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}

type errorBuilder struct{}

var unprocessableEntityError = &echo.HTTPError{
	Code: http.StatusUnprocessableEntity,
}
var badRequestError = &echo.HTTPError{
	Code:    http.StatusBadRequest,
	Message: "Unsupported Value Type",
}
var internalServerError = &echo.HTTPError{
	Code:    http.StatusInternalServerError,
	Message: "Internal server error",
}
var forbiddenError = &echo.HTTPError{
	Code: http.StatusForbidden,
}
var notFoundError = &echo.HTTPError{
	Code:    http.StatusNotFound,
	Message: "Resource not found",
}

func ErrorBuilder() *errorBuilder {
	return &errorBuilder{}
}

func (e *errorBuilder) NewFromDomain(ctx echo.Context, err errors.Error) error {
	if err.CausedByValidation() {
		return e.unprocessableEntityErrorWithMessage(ctx, err.String())
	} else if err.CausedInternally() {
		return ctx.JSON(internalServerError.Code, internalServerError)
	}
	badRequestError.Message = err.String()
	return ctx.JSON(badRequestError.Code, badRequestError)
}

func (*errorBuilder) NewForbiddenError() *echo.HTTPError {
	return forbiddenError
}


func (*errorBuilder) NewInternalServerError() *echo.HTTPError {
	return internalServerError
}

func (e *errorBuilder) NewNotFoundErrorWithMessage(ctx echo.Context, message string) error {
	notFoundError.Message = message
	return ctx.JSON(notFoundError.Code, notFoundError)
}

func (e *errorBuilder) NewForbiddenErrorWithMessage(ctx echo.Context, message string) error {
	forbiddenError.Message = message
	return ctx.JSON(forbiddenError.Code, forbiddenError)
}

func (*errorBuilder) badRequestErrorWithMessage(message string) *echo.HTTPError {
	err := badRequestError
	err.Message = message
	return err
}

func (*errorBuilder) internalErrorWithMessage(message string) *echo.HTTPError {
	err := internalServerError
	err.Message = message
	return err
}

func (*errorBuilder) unprocessableEntityErrorWithMessage(ctx echo.Context, message string) error {
	unprocessableEntityError.Message = message
	return ctx.JSON(unprocessableEntityError.Code, unprocessableEntityError)
}

func (e *ErrorMessage) Error() echo.HTTPError {
	return echo.HTTPError{
		Message: e.Message,
		Code:    e.Code,
	}
}

func (e *ErrorMessage) IsInternal() bool {
	return e.isInternal
}