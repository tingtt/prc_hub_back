package event

import (
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
)

type (
	CreateEventParam   = event.CreateEventParam
	UpdateEventParam   = event.UpdateEventParam
	GetEventQueryParam = event.GetEventQueryParam
)

func CreateEvent(p CreateEventParam, requestUserId string) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.CreateEvent(eventRepository, p, u)
}

func GetEvent(id string, requestUserId string) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.GetEvent(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		u,
	)
}

func GetEventList(q GetEventQueryParam, requestUserId string) (events []event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.GetEventList(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		event.GetEventQueryParam{
			Published:       q.Published,
			Name:            q.Name,
			NameContain:     q.NameContain,
			Location:        q.Location,
			LocationContain: q.LocationContain,
		},
		u,
	)
}

func UpdateEvent(id string, p UpdateEventParam, requestUserId string) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.UpdateEvent(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		p,
		u,
	)
}

func DeleteEvent(id string, requestUserId string) error {
	if !initialized {
		return ErrRepositoryNotInitialized
	}

	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return err
	}

	return event.DeleteEvent(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		u,
	)
}
