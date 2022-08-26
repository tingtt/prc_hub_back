package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (GET /events/{id})
func (s Server) GetEventsId(ctx echo.Context, id Id) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
