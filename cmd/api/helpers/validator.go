package helpers

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// CustomValidator is the validator implementation to be used with Echo.
type CustomValidator struct {
	validator *validator.Validate
}

var v = validator.New()

func init() {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tagName := fld.Tag.Get("json")
		if tagName == "" {

			tagName = fld.Tag.Get("form")
			if tagName == "" {
				tagName = fld.Tag.Get("query")

			}

		}

		name := strings.SplitN(tagName, ",", 2)[0]

		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
func FormatValidationErr(err error) map[string]string {
	errs := err.(validator.ValidationErrors)

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, trans)
	var errsMap = make(map[string]string)
	for _, fieldError := range errs {
		errsMap[fieldError.Field()] = fieldError.Translate(trans)
	}

	return errsMap
}
