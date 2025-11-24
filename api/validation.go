package api

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func ValidationError(c *gin.Context, err error) {
	var verr validator.ValidationErrors
	if !errors.As(err, &verr) {
		c.Error(errors.New("failed to parse validation error"))
		return
	}

	errs := parseValidationErrors(verr)

	c.JSON(http.StatusBadRequest, errs)
}

func parseValidationErrors(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, field := range verr {
		errs[field.Field()] = field.Tag()
	}

	return errs
}
