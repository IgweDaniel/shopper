package middleware

import (
	"reflect"

	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal"
	"github.com/labstack/echo/v4"
)

func ValidateDTO(dtoType interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Bind and validate the request body
			dto := reflect.New(reflect.TypeOf(dtoType).Elem()).Interface()
			if err := c.Bind(dto); err != nil {
				return helpers.HandleError(c, internal.WrapErrorMessage(internal.ErrBadRequest, "invalid JSON"))
			}

			if err := c.Validate(dto); err != nil {
				return helpers.HandleError(c, internal.WrapErrorMessage(internal.ErrBadRequest, "bad request"), helpers.FormatValidationErr(err))
			}

			// Store the validated DTO in the context
			c.Set("validatedDTO", dto)

			// Call the next handler
			if err := next(c); err != nil {
				c.Error(err)
			}

			// Remove the DTO from the context
			c.Set("validatedDTO", nil)

			return nil
		}
	}
}
