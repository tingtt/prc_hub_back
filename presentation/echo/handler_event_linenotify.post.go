package echo

import (
	"net/http"
	"prc_hub_back/application/event"
	"prc_hub_back/application/webhook"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
)

// (POST /events/{id}/webhook/line_notify)
func (s Server) PostEventsIdWebhookLineNotify(ctx echo.Context, id Id, params PostEventsIdWebhookLineNotifyParams) error {
	// Get jwt claim
	jcc, err := jwt.CheckProvided(ctx)
	if err != nil {
		return JSONMessage(ctx, http.StatusUnauthorized, err.Error())
	}

	// `Event`を取得
	e, err := event.GetEvent(
		id,
		event.GetEventQueryParam{
			Embed: &[]string{"documents"},
		},
		&jcc.Id,
	)
	if err != nil {
		return JSONMessage(ctx, event.ErrToCode(err), err.Error())
	}

	// Scope
	if params.Scope == nil {
		var defaultScope PostEventsIdWebhookLineNotifyParamsScope = PostEventsIdWebhookLineNotifyParamsScope(LineNotifyScopeEvent)
		params.Scope = &defaultScope
	}
	switch *params.Scope {
	case PostEventsIdWebhookLineNotifyParamsScope(LineNotifyScopeDocument):
		// notify documents to line
		err = webhook.NotifyEventDocuments(e)
	default:
		// notify event to line
		err = webhook.NotifyEvent(e)
	}
	if err != nil {
		return JSONMessage(ctx, http.StatusInternalServerError, err.Error())
	}

	return JSONMessage(ctx, http.StatusOK, "ok")
}
