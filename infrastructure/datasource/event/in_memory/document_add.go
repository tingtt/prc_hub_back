package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEventDocument) Add(e event.EventDocument) (event.EventDocument, error) {
	dataDocument = append(dataDocument, e)
	return e, nil
}
