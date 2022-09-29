package event_inmemory

import (
	"prc_hub_back/domain/model/event"
	"time"
)

func (r RepositoryEvent) Update(id string, p event.UpdateEventParam) (_ event.Event, err error) {
	var e event.Event
	err = event.ErrEventNotFound
	for _, e2 := range dataEvent {
		if e2.Id == id {
			err = nil
			e = e2
		}
	}
	if err != nil {
		return
	}

	if p.Name != nil {
		e.Name = *p.Name
	}
	if p.Description.KeyExists() {
		if p.Description.IsNull() {
			e.Description = nil
		} else {
			e.Description = *p.Description.Value
		}
	}
	if p.Location.KeyExists() {
		if p.Location.IsNull() {
			e.Location = nil
		} else {
			e.Location = *p.Location.Value
		}
	}
	if p.Datetimes != nil {
		for _, d := range *p.Datetimes {
			e.Datetimes = append(e.Datetimes, event.EventDatetime{
				Start: time.Time(d.Start),
				End:   time.Time(d.End),
			})
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
