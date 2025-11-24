package api

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func RegisterValidation() ut.Translator {

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return trans
}

func ValidationError(c *gin.Context, err error) {
	trans := GetContextTranslation(c)

	var verr validator.ValidationErrors
	if !errors.As(err, &verr) {
		c.Error(errors.New("failed to parse validation error"))
		return
	}

	errs := make(map[string]string)

	for _, field := range verr {
		errs[field.Field()] = field.Translate(trans)
	}

	c.JSON(http.StatusBadRequest, errs)
}
