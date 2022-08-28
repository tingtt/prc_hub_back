package echo

import (
	"net/http"
	"prc_hub_back/application/event"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
)

// (GET /events/{id})
func (s Server) GetEventsId(ctx echo.Context, id Id) error {
	// Get jwt claim
	jcc, err := jwt.Check(ctx)
	if err != nil {
		return JSONMessage(ctx, http.StatusUnauthorized, err.Error())
	}

	// Get event
	e, err := event.GetEvent(id, jcc.Id)
	if err != nil {
		return JSONMessage(ctx, event.ErrToCode(err), err.Error())
	}

	return JSONPretty(ctx, http.StatusOK, e)
}
