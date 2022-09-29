package event

import (
	"errors"
	"prc_hub_back/domain/model/user"
)

// Errors
var (
	ErrCannotUpdateEventDocument = errors.New("sorry, you cannot update this document")
)

type UpdateEventDocumentParam struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

func (p UpdateEventDocumentParam) validate(repo EventRepository, qs EventQueryService, id string, requestUser user.User) error {
	/**
	 * フィールドの検証
	**/
	if p.Name == nil && p.Url == nil {
		return ErrNoUpdates
	}
	// `Name`
	if p.Name != nil {
		err := validateDocumentName(*p.Name)
		if err != nil {
			return err
		}
	}
	// `Url`
	if p.Url != nil {
		err := validateUrl(*p.Url)
		if err != nil {
			return err
		}
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage {
		ed, err := GetDocument(repo, qs, id, requestUser)
		if err != nil {
			return err
		}

		// Eventを取得
		e, err := GetEvent(qs, ed.EventId, GetEventQueryParam{}, requestUser)
		if err != nil {
			return err
		}

		if e.UserId != requestUser.Id {
			// `User`が`Admin`・`Manage`のいずれでもなく
			// `Published`でない 且つ 自分のものでない`Event`は変更不可
			return ErrCannotUpdateEventDocument
		}
	}

	return nil
}

func UpdateEventDocument(repo EventRepository, qs EventQueryService, id string, p UpdateEventDocumentParam, requestUser user.User) (_ EventDocument, err error) {
	// 確認
	_, err = repo.GetDocument(id)
	if err != nil {
		return
	}

	// バリデーション
	err = p.validate(repo, qs, id, requestUser)
	if err != nil {
		return
	}

	return repo.UpdateDocument(id, p)
}
