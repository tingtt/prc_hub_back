package event

import "errors"

// Errors
var (
	ErrEventNotFound         = errors.New("event not found")
	ErrEventDocumentNotFound = errors.New("event document not found")
	ErrNoUpdates             = errors.New("no updates")
)

type EventRepository interface {
	Add(e Event) (Event, error)
	Update(id string, p UpdateEventParam) (Event, error)
	Delete(id string) error
	AddDocument(e EventDocument) (EventDocument, error)
	GetDocument(id string) (EventDocument, error)
	GetDocumentList(q GetDocumentQueryParam) ([]EventDocument, error)
	UpdateDocument(id string, p UpdateEventDocumentParam) (EventDocument, error)
	DeleteDocument(id string) error
}

type EventQueryService interface {
	Get(id string, q GetEventQueryParam) (EventEmbed, error)
	GetList(q GetEventListQueryParam) ([]EventEmbed, error)
}
