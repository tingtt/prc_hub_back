package event_inmemory

import (
	eventApp "prc_hub_back/application/event"
	userApp "prc_hub_back/application/user"
	"prc_hub_back/domain/model/event"
	"prc_hub_back/domain/model/user"
)

func (r RepositoryEvent) Get(id string, q event.GetEventQueryParam) (_ event.Event, err error) {
	for _, e := range dataEvent {
		if e.Id == id {
			if q.Embed != nil {
				// 埋め込みフィールドの取得
				for _, embed := range *q.Embed {
					if embed == "user" {
						// `User`を取得
						var u user.User
						u, err = userApp.Get(e.UserId)
						if err != nil && err != userApp.ErrUserNotFound {
							return
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
						var u user.User
						u, err = userApp.Get(e.UserId)
						if err != nil && err != userApp.ErrUserNotFound {
							return
						}
						if err == userApp.ErrUserNotFound {
							// `User`が存在しない場合
							u = user.User{Name: "Deleted user"}
						}
						// `Document`を取得
						var documents []event.EventDocument
						documents, err = eventApp.GetDocumentList(
							eventApp.GetDocumentQueryParam{
								EventId: &e.Id,
							},
							u.Id,
						)
						if err != nil {
							return
						}
						// `Document`を埋め込み
						e.Documents = &documents
					}
				}
			}
			return e, nil
		}
	}
	err = event.ErrEventNotFound
	return
}
