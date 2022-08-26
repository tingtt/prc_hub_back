package event_inmemory

import "prc_hub_back/domain/model/event"

type (
	RepositoryEvent         struct{}
	RepositoryEventDocument struct{}
)

var (
	dataEvent    []event.Event
	dataDocument []event.EventDocument
)
