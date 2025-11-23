package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) {

	// Healthcheck
	router.GET("/healthcheck", healthcheck)

	router.Use(ErrorMiddleware)
	router.Use(TransactionMiddleware(db))
}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
