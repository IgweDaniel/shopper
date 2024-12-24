package helpers

import (
	"net/http"

	"github.com/IgweDaniel/shopper/internal"
	"github.com/labstack/echo/v4"
)

func HandleError(ctx echo.Context, err error, customErrorResponse ...interface{}) *echo.HTTPError {
	var context string

	if wrappedErr, ok := err.(*internal.WrappedError); ok {
		context = wrappedErr.Context
		err = wrappedErr.Err
	} else {
		ctx.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, buildErrorResponse("internal server error", nil))
	}
	// err = errors.Unwrap(err)
	var statusCode int
	switch err {
	case internal.ErrNotFound:
		statusCode = http.StatusNotFound
	case internal.ErrNotAuthorized:
		statusCode = http.StatusForbidden
	case internal.ErrBadRequest:
		statusCode = http.StatusBadRequest
	case internal.ErrDuplicatedKey:
		statusCode = http.StatusConflict
	case internal.ErrRateLimit:
		statusCode = http.StatusTooManyRequests
	case internal.ErrInternal:
		ctx.Logger().Error(err)
		statusCode = http.StatusInternalServerError
	default:
		ctx.Logger().Error(err)
		statusCode = http.StatusInternalServerError
	}

	if len(customErrorResponse) > 0 {
		return echo.NewHTTPError(statusCode, buildErrorResponse(context, customErrorResponse[0]))
	}

	return echo.NewHTTPError(statusCode, buildErrorResponse(context, nil))
}

func buildErrorResponse(respMessage string, respData interface{}) echo.Map {
	return buildResponse(false, respMessage, respData)
}

func BuildResponse(respMessage string, respData interface{}) echo.Map {
	return buildResponse(true, respMessage, respData)
}

func buildResponse(success bool, respMessage string, respData interface{}) echo.Map {
	return echo.Map{
		"success": success,
		"message": respMessage,
		"data":    respData,
	}
}
