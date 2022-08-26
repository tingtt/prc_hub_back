package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// (POST /events/{id}/webhook/line_notify)
func (s Server) PostEventsIdWebhookLineNotify(ctx echo.Context, id Id, params PostEventsIdWebhookLineNotifyParams) error {
	if params.Scope == nil {
		var defaultScope PostEventsIdWebhookLineNotifyParamsScope = PostEventsIdWebhookLineNotifyParamsScope(LineNotifyScopeEvent)
		params.Scope = &defaultScope
	}
	switch *params.Scope {
	case PostEventsIdWebhookLineNotifyParamsScope(LineNotifyScopeDocument):
		// TODO: notify document to line
	default:
		// TODO: notify event to line
	}

	// TODO
	return JSONMessage(ctx, http.StatusInternalServerError, "not inplemented")
}
