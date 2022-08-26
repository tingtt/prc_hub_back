package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) Add(e event.Event) (event.Event, error) {
	dataEvent = append(dataEvent, e)
	return e, nil
}
