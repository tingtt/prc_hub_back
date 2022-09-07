package event_inmemory

import (
	eventApp "prc_hub_back/application/event"
	userApp "prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
	"prc_hub_back/domain/model/user"
)

func (r RepositoryEvent) GetList(q event.GetEventListQueryParam) (_ []event.Event, err error) {
	if q.Name == nil &&
		q.NameContain == nil &&
		q.Location == nil &&
		q.LocationContain == nil &&
		q.Published == nil &&
		q.Embed == nil {
		return dataEvent, nil
	}

	var events []event.Event
	for _, e := range dataEvent {
		if q.Embed != nil {
			// 埋め込みフィールドの取得
			for _, embed := range *q.Embed {
				if embed == "user" {
					// `User`を取得
					u, err := userApp.Get(e.UserId)
					if err != nil && err != userApp.ErrUserNotFound {
						return nil, err
					}
					if err == userApp.ErrUserNotFound {
						// `User`が存在しない場合
						u = user.User{Name: "Deleted user"}
					}
					// `User`を埋め込み
					e.User = &u
				}
				if embed == "documents" {
					// `User`を取得
					u, err := userApp.Get(e.UserId)
					if err != nil && err != userApp.ErrUserNotFound {
						return nil, err
					}
					if err == userApp.ErrUserNotFound {
						// `User`が存在しない場合
						u = user.User{Name: "Deleted user"}
					}
					// `Document`を取得
					documents, err := eventApp.GetDocumentList(
						eventApp.GetDocumentQueryParam{
							EventId: &e.Id,
						},
						u.Id,
					)
					if err != nil {
						return nil, err
					}
					// `Document`を埋め込み
					e.Documents = &documents
				}
			}
		}
		events = append(events, e)
	}
	return events, nil
}
