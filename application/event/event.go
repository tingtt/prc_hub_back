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
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.CreateEvent(p, u)
}

func GetEvent(id string, q GetEventQueryParam, requestUserId *string) (_ event.EventEmbed, err error) {
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
		id,
		q,
		*u,
	)
}

func GetEventList(q GetEventListQueryParam, requestUserId *string) (events []event.EventEmbed, err error) {
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
		q,
		*u,
	)
}

func UpdateEvent(id string, p UpdateEventParam, requestUserId string) (_ event.Event, err error) {
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return
	}

	return event.UpdateEvent(
		id,
		p,
		u,
	)
}

func DeleteEvent(id string, requestUserId string) error {
	// リクエスト元のユーザーを取得
	u, err := user.Get(requestUserId)
	if err != nil {
		return err
	}

	return event.DeleteEvent(
		id,
		u,
	)
}
