package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) GetList(q event.GetEventQueryParam) (_ []event.Event, err error) {
	return dataEvent, nil
}
