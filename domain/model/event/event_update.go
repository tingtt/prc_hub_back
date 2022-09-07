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
	Name        *string                     `json:"name,omitempty"`
	Description util.NullableJSONString     `json:"description,omitempty"`
	Location    util.NullableJSONString     `json:"location,omitempty"`
	Datetimes   *[]CreateEventDatetimeParam `json:"datetimes,omitempty"`
	Published   *bool                       `json:"published,omitempty"`
	Completed   *bool                       `json:"completed,omitempty"`
}

func (p UpdateEventParam) validate(repo Repos, id string, requestUser user.User) error {
	/**
	 * フィールドの検証
	**/
	if p.Name == nil &&
		p.Description.KeyExists() &&
		p.Location.KeyExists() &&
		p.Datetimes == nil &&
		p.Published == nil &&
		p.Completed == nil {
		return ErrNoUpdates
	}
	// `Name`
	if p.Name != nil {
		err := validateTitle(*p.Name)
		if err != nil {
			return err
		}
	}
	// `Datetimes`
	if p.Datetimes != nil {
		if len(*p.Datetimes) == 0 {
			return ErrValidateEventDatetimesCannotBeEmpty
		}
		for _, d := range *p.Datetimes {
			err := d.validate()
			if err != nil {
				return err
			}
		}
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		// Eventを取得
		e, err := GetEvent(repo, id, GetEventQueryParam{}, requestUser)
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
