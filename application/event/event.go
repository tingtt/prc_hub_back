package event

import (
	"prc_hub_back/domain/model/event"
)

func Create(p event.CreateEventParam) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.CreateEvent(eventRepository, p)
}

func Get(id string) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.GetEvent(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
	)
}

func Update(id string, p event.UpdateEventParam) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.UpdateEvent(eventRepository, id, p)
}

func Delete(id string) error {
	if !initialized {
		return ErrRepositoryNotInitialized
	}
	return event.DeleteEvent(eventRepository, id)
}
