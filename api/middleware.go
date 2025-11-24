package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
)

func TransactionMiddleware(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := DB.Begin()
		SetContextTransaction(c, tx)

		c.Next()

		if len(c.Errors) == 0 {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}
}

func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error": c.Errors.Last().Err,
		})
	}
}

func TranslationMiddleware(trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		SetContextTranslation(c, trans)

		c.Next()
	}
}
