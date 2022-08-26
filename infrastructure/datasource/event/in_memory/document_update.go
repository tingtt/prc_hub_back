package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEventDocument) Update(id string, p event.UpdateEventDocumentParam) (_ event.EventDocument, err error) {
	ed, err := r.Get(id)
	if err != nil {
		return
	}

	if p.Name != nil {
		ed.Name = *p.Name
	}
	if p.Url != nil {
		ed.Url = *p.Url
	}

	err = r.Delete(id)
	if err != nil {
		return
	}
	_, err = r.Add(ed)
	if err != nil {
		return
	}

	return ed, nil
}
