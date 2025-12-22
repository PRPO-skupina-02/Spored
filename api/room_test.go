package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PRPO-skupina-02/common/database"
	"github.com/PRPO-skupina-02/common/xtesting"
	"github.com/orgs/PRPO-skupina-02/Spored/db"
	"github.com/stretchr/testify/assert"
)

func TestRoomsList(t *testing.T) {
	db, fixtures := database.PrepareTestDatabase(t, db.FixtureFS, db.MigrationsFS)
	r := TestingRouter(t, db)

	tests := []struct {
		name      string
		status    int
		params    string
		theaterID string
	}{
		{
			name:      "ok",
			status:    http.StatusOK,
			theaterID: "fb126c8c-d059-11f0-8fa4-b35f33be83b7",
		},
		{
			name:      "ok-paginated",
			status:    http.StatusOK,
			params:    "?limit=1&offset=1",
			theaterID: "bae209f6-d059-11f0-b2a4-cbf992c2eb6d",
		},
		{
			name:      "ok-sort",
			status:    http.StatusOK,
			params:    "?sort=-updated_at",
			theaterID: "bae209f6-d059-11f0-b2a4-cbf992c2eb6d",
		},
		{
			name:      "ok-paginated-sort",
			status:    http.StatusOK,
			params:    "?limit=2&offset=1&sort=updated_at",
			theaterID: "bae209f6-d059-11f0-b2a4-cbf992c2eb6d",
		},
		{
			name:      "ok-no-rooms",
			status:    http.StatusOK,
			theaterID: "ea0b7f96-ddc9-11f0-9635-23efd36396bd",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := fixtures.Load()
			assert.NoError(t, err)

			targetURL := fmt.Sprintf("/api/v1/theaters/%s/rooms%s", testCase.theaterID, testCase.params)

			req := xtesting.NewTestingRequest(t, targetURL, http.MethodGet, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.status, w.Code)
			xtesting.AssertGoldenJSON(t, w)
		})
	}
}
