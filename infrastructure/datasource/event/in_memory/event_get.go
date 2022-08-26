package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) Get(id string) (_ event.Event, err error) {
	for _, e := range dataEvent {
		if e.Id == id {
			return e, nil
		}
	}
	err = event.ErrEventNotFound
	return
}
