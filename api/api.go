package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) {

	RegisterValidation()

	// Healthcheck
	router.GET("/healthcheck", healthcheck)

	api := router.Group("/api")
	api.Use(ErrorMiddleware)
	api.Use(TransactionMiddleware(db))

	// Theaters
	theaters := api.Group("/theaters")
	theaters.GET("/", TheatersList)
	theaters.POST("/", TheatersCreate)

}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
