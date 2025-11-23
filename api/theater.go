package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/orgs/PRPO-skupina-02/Spored/models"
)

type TheaterResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func newTheaterResponse(theater models.Theater) TheaterResponse {
	return TheaterResponse{
		UUID:      theater.UUID,
		CreatedAt: theater.CreatedAt,
		UpdatedAt: theater.UpdatedAt,
		Name:      theater.Name,
	}
}

func TheatersList(c *gin.Context) {
	tx := GetContextTransaction(c)

	var theaters []models.Theater
	if err := tx.Find(&theaters).Error; err != nil {
		_ = c.Error(err)
		return
	}

	response := []TheaterResponse{}

	for _, theater := range theaters {
		response = append(response, newTheaterResponse(theater))
	}

	c.JSON(http.StatusOK, response)
}
