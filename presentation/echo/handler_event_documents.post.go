package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (POST /events/{id}/documents)
func (s Server) PostEventsIdDocuments(ctx echo.Context, id Id) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
