package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
	"prc_hub_back/domain/model/util"
)

// Errors
var (
	ErrCannotCreateEventDocument = errors.New("sorry, you cannot create document to this event")
)

type CreateEventDocumentParam struct {
	EventId string `json:"event_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

func (p CreateEventDocumentParam) validate(repo EventDocumentRepository, qs EventQueryService, requestUser user.User) error {
	// フィールドの検証
	err := validateDocumentName(p.Name)
	if err != nil {
		return err
	}
	err = validateUrl(p.Url)
	if err != nil {
		return err
	}
	err = validateEventId(qs, p.EventId)
	if err != nil {
		return err
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		// Eventを取得
		e, err := GetEvent(qs, p.EventId, GetEventQueryParam{}, requestUser)
		if err != nil {
			return err
		}

		if requestUser.Id != e.UserId {
			// `Admin`・`Manage`のいずれでもなく`Event.UserId`が自分ではない場合は追加不可
			return ErrCannotCreateEventDocument
		}
	}

	return nil
}

func CreateEventDocument(repo EventDocumentRepository, qs EventQueryService, p CreateEventDocumentParam, requestUser user.User) (_ EventDocument, err error) {
	err = p.validate(repo, qs, requestUser)
	if err != nil {
		return
	}

	return repo.Add(EventDocument{
		Id:      util.UUID(),
		EventId: p.EventId,
		Name:    p.Name,
		Url:     p.Url,
	})
}
