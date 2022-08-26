package event

import "errors"

// Errors
var (
	ErrEventNotFound         = errors.New("event not found")
	ErrEventDocumentNotFound = errors.New("event document not found")
)

type EventRepository interface {
	Add(e Event) (Event, error)
	Get(id string) (Event, error)
	GetList(q GetEventQueryParam) ([]Event, error)
	Update(id string, p UpdateEventParam) (Event, error)
	Delete(id string) error
}

type EventDocumentRepository interface {
	Add(e EventDocument) (EventDocument, error)
	Get(id string) (EventDocument, error)
	GetList(q GetDocumentQueryParam) ([]EventDocument, error)
	Update(id string, p UpdateEventDocumentParam) (EventDocument, error)
	Delete(id string) error
}

type Repos struct {
	Event    EventRepository
	Document EventDocumentRepository
}
