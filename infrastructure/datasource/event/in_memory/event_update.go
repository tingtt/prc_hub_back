package event_inmemory

import "prc_hub_back/domain/model/event"

func (r RepositoryEvent) Update(id string, p event.UpdateEventParam) (_ event.Event, err error) {
	e, err := r.Get(id)
	if err != nil {
		return
	}

	if p.Name != nil {
		e.Name = *p.Name
	}
	if p.Location.KeyExists() {
		if p.Location.IsNull() {
			e.Location = nil
		} else {
			e.Location = *p.Location.Value
		}
	}
	if p.Published != nil {
		e.Published = *p.Published
	}
	if p.Completed != nil {
		e.Completed = *p.Completed
	}

	err = r.Delete(id)
	if err != nil {
		return
	}
	_, err = r.Add(e)
	if err != nil {
		return
	}

	return e, nil
}
