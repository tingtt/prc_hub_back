package event

import (
	"errors"
	"prc_hub_back/domain/model/event"
)

var (
	initialized        = false
	eventRepository    event.EventRepository
	documentRepository event.EventDocumentRepository
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(eRepo event.EventRepository, dRepo event.EventDocumentRepository) {
	initialized = true
	eventRepository = eRepo
	documentRepository = dRepo
}
