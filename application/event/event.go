package event

import (
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
	userDomain "prc_hub_back/domain/model/user"
)

type (
	CreateEventParam       = event.CreateEventParam
	UpdateEventParam       = event.UpdateEventParam
	GetEventListQueryParam = event.GetEventListQueryParam
	GetEventQueryParam     = event.GetEventQueryParam
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

func GetEvent(id string, q GetEventQueryParam, requestUserId *string) (_ event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	u := new(userDomain.User)

	if requestUserId != nil {
		// リクエスト元のユーザーを取得
		var u2 userDomain.User
		u2, err = user.Get(*requestUserId)
		if err != nil {
			return
		}
		u = &u2
	} else if requestUserId == nil {
		// リクエストユーザーが指定されていない場合は最小権限のユーザーを仮使用
		u = &userDomain.User{
			Id:                  "",
			PostEventAvailabled: false,
			Manage:              false,
			Admin:               false,
		}
	}

	return event.GetEvent(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		id,
		q,
		*u,
	)
}

func GetEventList(q GetEventListQueryParam, requestUserId *string) (events []event.Event, err error) {
	if !initialized {
		err = ErrRepositoryNotInitialized
		return
	}

	u := new(userDomain.User)

	if requestUserId != nil {
		// リクエスト元のユーザーを取得
		var u2 userDomain.User
		u2, err = user.Get(*requestUserId)
		if err != nil {
			return
		}
		u = &u2
	} else if requestUserId == nil {
		// リクエストユーザーが指定されていない場合は最小権限のユーザーを仮使用
		u = &userDomain.User{
			Id:                  "",
			PostEventAvailabled: false,
			Manage:              false,
			Admin:               false,
		}
	}

	return event.GetEventList(
		event.Repos{
			Event:    eventRepository,
			Document: documentRepository,
		},
		event.GetEventListQueryParam{
			Published:       q.Published,
			Name:            q.Name,
			NameContain:     q.NameContain,
			Location:        q.Location,
			LocationContain: q.LocationContain,
		},
		*u,
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
