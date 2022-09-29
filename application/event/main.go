package event

import (
	"errors"
	"prc_hub_back/domain/model/event"
)

// Singleton fields
var (
	initialized  = false
	repo         event.EventRepository
	queryService event.EventQueryService
)

// Errors
var (
	ErrRepositoryNotInitialized = errors.New("repository not initialized")
)

func InitApplication(eRepo event.EventRepository, eQueryService event.EventQueryService) {
	initialized = true
	repo = eRepo
	queryService = eQueryService
}
