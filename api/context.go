package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const contextTransactionKey = "transaction"

func SetContextTransaction(c *gin.Context, tx *gorm.DB) {
	c.Set(contextTransactionKey, tx)
}

func GetContextTransaction(c *gin.Context) *gorm.DB {
	tx, ok := c.Get(contextTransactionKey)
	if !ok {
		slog.Error("Could not get transaction from context")
		return nil
	}

	return tx.(*gorm.DB)
}
