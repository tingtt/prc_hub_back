package echo

import (
	"net/http"
	"prc_hub_back/application/event"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
)

// (GET /events/{id}/documents/{document_id})
func (s Server) GetEventsIdDocumentsDocumentId(ctx echo.Context, id Id, documentId DocumentId) error {
	// Get jwt claim
	jcc, err := jwt.Check(ctx)
	if err != nil {
		return JSONMessage(ctx, http.StatusUnauthorized, err.Error())
	}

	// Get document
	ed, err := event.GetDocument(documentId, jcc.Id)
	if err != nil {
		return JSONMessage(ctx, event.ErrToCode(err), err.Error())
	}

	return JSONPretty(ctx, http.StatusOK, ed)
}
