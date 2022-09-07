package event

import "prc_hub_back/domain/model/user"

type GetEventQueryParam struct {
	Embed *[]string `query:"embed"`
}

func GetEvent(repo Repos, id string, q GetEventQueryParam, requestUser user.User) (e Event, err error) {
	// Get event
	tmpE, err := repo.Event.Get(id, q)
	if err != nil {
		return
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage &&
		!tmpE.Published && tmpE.UserId != requestUser.Id {
		// `User`が`Admin`・`Manage`のいずれでもなく
		// `Published`でない 且つ 自分のものでない`Event`は取得不可
		err = ErrEventNotFound
		return
	}

	e = tmpE
	return
}
