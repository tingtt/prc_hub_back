package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) Delete(id string) (err error) {
	old := dataEvent
	dataEvent = nil
	err = event.ErrEventNotFound
	for _, d := range old {
		if d.Id == id {
			err = nil
		} else {
			dataEvent = append(dataEvent, d)
		}
	}
	return
}
