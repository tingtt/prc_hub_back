package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (GET /events)
func (s Server) GetEvents(ctx echo.Context, params GetEventsParams) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
