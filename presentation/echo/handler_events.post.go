package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (POST /events)
func (s Server) PostEvents(ctx echo.Context) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
