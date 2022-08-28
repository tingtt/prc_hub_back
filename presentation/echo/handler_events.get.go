package echo

import (
	"net/http"
	"prc_hub_back/application/event"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
)

// (GET /events)
func (s Server) GetEvents(ctx echo.Context, params GetEventsParams) error {
	// Get jwt claim
	jcc, err := jwt.Check(ctx)
	if err != nil {
		return JSONMessage(ctx, http.StatusUnauthorized, err.Error())
	}

	// Bind query
	query := new(event.GetEventQueryParam)
	if err := ctx.Bind(query); err != nil {
		return JSONMessage(ctx, http.StatusBadRequest, err.Error())
	}

	// Get events
	events, err := event.GetEventList(*query, jcc.Id)
	if err != nil {
		return JSONMessage(ctx, event.ErrToCode(err), err.Error())
	}

	return JSONPretty(ctx, http.StatusOK, events)
}
