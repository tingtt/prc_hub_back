package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
	"prc_hub_back/domain/model/util"
)

// Errors
var (
	ErrCannotCreateEvent = errors.New("sorry, you cannot create `event`")
)

type CreateEventParam struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Location    *string `json:"location,omitempty"`
	Published   bool    `json:"published"`
	Completed   bool    `json:"completed"`
}

func (p CreateEventParam) validate(requestUser user.User) error {
	// フィールドの検証
	err := validateTitle(p.Name)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage && !requestUser.PostEventAvailabled {
		// `Admin`・`Manage`・`PostEventAvailabled`のいずれでもない場合は`Event`作成不可
		return ErrCannotCreateEvent
	}

	return nil
}

func CreateEvent(repo EventRepository, p CreateEventParam, requestUser user.User) (_ Event, err error) {
	// バリデーション
	err = p.validate(requestUser)
	if err != nil {
		return
	}

	return repo.Add(Event{
		Id:          util.UUID(),
		Name:        p.Name,
		Description: p.Description,
		Location:    p.Location,
		Published:   p.Published,
		Completed:   p.Completed,
		UserId:      requestUser.Id,
	})
}
