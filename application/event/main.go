package event

import (
	"errors"
	"prc_hub_back/domain/model/event"
)

// Singleton fields
var (
	initialized        = false
	eventRepository    event.EventRepository
	documentRepository event.EventDocumentRepository
	eventQueryService  event.EventQueryService
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(eRepo event.EventRepository, dRepo event.EventDocumentRepository, eQueryService event.EventQueryService) {
	initialized = true
	eventRepository = eRepo
	documentRepository = dRepo
	eventQueryService = eQueryService
}
