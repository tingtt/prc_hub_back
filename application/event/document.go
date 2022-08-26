package event

import (
	"prc_hub_back/domain/model/event"
)

func CreateDocument(p event.CreateEventDocumentParam) (_ event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.CreateEventDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		p,
	)
}

func GetDocument(id string) (_ event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.GetDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
	)
}

func UpdateDocument(id string, p event.UpdateEventDocumentParam) (_ event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.UpdateEventDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		p,
	)
}

func DeleteDocument(id string) error {
	if !initialized {
		return ErrRepositoryNotInitialized
	}
	return event.DeleteEventDocument(documentRepository, id)
}
