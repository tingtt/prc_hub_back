package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (PUT /events/{id}/documents)
func (s Server) PutEventsIdDocuments(ctx echo.Context, id Id) error {
	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
