package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEventDocument) GetList(q event.GetDocumentQueryParam) (_ []event.EventDocument, err error) {
	return dataDocument, nil
}
