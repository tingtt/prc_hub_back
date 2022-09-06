package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
	"prc_hub_back/domain/model/util"
)

// Errors
var (
	ErrCannotUpdateEvent = errors.New("sorry, you cannot update this event")
)

type UpdateEventParam struct {
	Name        *string                 `json:"name"`
	Description util.NullableJSONString `json:"description,omitempty"`
	Location    util.NullableJSONString `json:"location,omitempty"`
	Published   *bool                   `json:"published"`
	Completed   *bool                   `json:"completed"`
}

func (p UpdateEventParam) validate(repo Repos, id string, requestUser user.User) error {
	// フィールドの検証
	if p.Name != nil {
		err := validateTitle(*p.Name)
		if err != nil {
			return err
		}
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		// Eventを取得
		e, err := GetEvent(repo, id, requestUser)
		if err != nil {
			return err
		}

		if requestUser.Id != e.UserId {
			// `Admin`・`Manage`のいずれでもなく`Event.UserId`が自分ではない場合は変更不可
			return ErrCannotUpdateEvent
		}
	}

	return nil
}

func UpdateEvent(repo Repos, id string, p UpdateEventParam, requestUser user.User) (_ Event, err error) {
	// バリデーション
	err = p.validate(repo, id, requestUser)
	if err != nil {
		return
	}

	return repo.Event.Update(id, p)
}
