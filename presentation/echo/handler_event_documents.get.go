package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (GET /events/{id}/documents)
func (s Server) GetEventsIdDocuments(ctx echo.Context, id Id, params GetEventsIdDocumentsParams) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
