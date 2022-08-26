package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (PATCH /events/{id})
func (s Server) PatchEventsId(ctx echo.Context, id Id) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
