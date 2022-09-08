package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
	"prc_hub_back/domain/model/util"
	"time"
)

// Errors
var (
	ErrCannotCreateEvent = errors.New("sorry, you cannot create `event`")
)

type CreateEventParam struct {
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	Datetimes   []CreateEventDatetimeParam `json:"datetimes"`
	Location    *string                    `json:"location,omitempty"`
	Published   bool                       `json:"published"`
	Completed   bool                       `json:"completed"`
}

type CreateEventDatetimeParam struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (p CreateEventDatetimeParam) validate() error {
	// フィールドの検証
	err := validateEventDatetime(time.Time(p.Start), time.Time(p.End))
	if err != nil {
		return err
	}
	return nil
}

func (p CreateEventParam) validate(requestUser user.User) error {
	/**
	 * フィールドの検証
	**/
	// `Name`
	err := validateTitle(p.Name)
	if err != nil {
		return err
	}
	// `Datetimes`
	if len(p.Datetimes) == 0 {
		return ErrValidateDocumentNameCannotBeEmpty
	}
	for _, d := range p.Datetimes {
		err = d.validate()
		if err != nil {
			return err
		}
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

	var datetimes []EventDatetime
	for _, d := range p.Datetimes {
		datetimes = append(datetimes, EventDatetime{
			Start: d.Start.UTC(),
			End:   d.End.UTC(),
		})
	}

	return repo.Add(Event{
		Id:          util.UUID(),
		Name:        p.Name,
		Description: p.Description,
		Location:    p.Location,
		Datetimes:   datetimes,
		Published:   p.Published,
		Completed:   p.Completed,
		UserId:      requestUser.Id,
	})
}
