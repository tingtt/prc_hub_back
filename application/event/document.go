package event

import (
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
)

type (
	CreateEventDocumentParam event.CreateEventDocumentParam
	UpdateEventDocumentParam event.UpdateEventDocumentParam
	GetDocumentQueryParam    event.GetDocumentQueryParam
)

func CreateDocument(p CreateEventDocumentParam, requestUserId string) (_ event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.CreateEventDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		event.CreateEventDocumentParam{
			EventId: p.EventId,
			Name:    p.Name,
			Url:     p.Url,
		},
		u,
	)
}

func GetDocument(id string, requestUserId string) (_ event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.GetDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		u,
	)
}

func GetDocumentList(q GetDocumentQueryParam, requestUserId string) (documents []event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}
	return event.GetDocumentList(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		event.GetDocumentQueryParam{
			EventId:     q.EventId,
			Name:        q.Name,
			NameContain: q.NameContain,
		},
	)
}

func UpdateDocument(id string, p UpdateEventDocumentParam, requestUserId string) (_ event.EventDocument, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.UpdateEventDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		event.UpdateEventDocumentParam{
			Name: p.Name,
			Url:  p.Url,
		},
		u,
	)
}

func DeleteDocument(id string, requestUserId string) error {
	if !initialized {
		return ErrRepositoryNotInitialized
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return err
	}

	return event.DeleteEventDocument(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		u,
	)
}
