package api

import (
	"time"

	"github.com/PRPO-skupina-02/common/middleware"
	"github.com/PRPO-skupina-02/common/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/orgs/PRPO-skupina-02/Spored/models"
)

type RoomResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func newRoomResponse(room models.Room) RoomResponse {
	return RoomResponse{
		ID:        room.ID,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
		Name:      room.Name,
	}
}

// RoomsList
//
//	@Id				RoomsList
//	@Summary		List theater rooms
//	@Description	List theater rooms
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int		false	"Limit the number of responses"	Default(10)
//	@Param			offset	query		int		false	"Offset the first response"		Default(0)
//	@Param			sort	query		string	false	"Sort results"
//	@Success		200		{object}	[]RoomResponse
//	@Failure		400		{object}	middleware.HttpError
//	@Failure		404		{object}	middleware.HttpError
//	@Failure		500		{object}	middleware.HttpError
//	@Router			/theaters/:theaterID/rooms [get]
func RoomsList(c *gin.Context) {
	tx := middleware.GetContextTransaction(c)
	offset, limit := request.GetNormalizedPaginationArgs(c)
	sort := request.GetSortOptions(c)
	theaterID, err := request.GetUUIDParam(c, "theaterID")
	if err != nil {
		_ = c.Error(err)
		return
	}

	rooms, total, err := models.GetTheaterRooms(tx, theaterID, offset, limit, sort)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response := []RoomResponse{}

	for _, room := range rooms {
		response = append(response, newRoomResponse(room))
	}

	request.RenderPaginatedResponse(c, response, total)
}
