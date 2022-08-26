package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (DELETE /events/{id})
func (s Server) DeleteEventsId(ctx echo.Context, id Id) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
