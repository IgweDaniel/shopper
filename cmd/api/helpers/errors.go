package helpers

import (
	"github.com/labstack/echo/v4"
)

// ErrInternalServer logs the passed in error and returns a 500 message.
func ErrInternalServer(ctx echo.Context, err error) error {
	ctx.Logger().Error(err)

	// by default, Echo returns standard errors as internal server errors
	// https://echo.labstack.com/guide/error-handling/#default-http-error-handler
	return err
}
