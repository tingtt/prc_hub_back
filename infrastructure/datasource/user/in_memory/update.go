package user_inmemory

import (
	"prc_hub_back/domain/model/user"
)

func (r Repository) Update(id string, p user.UpdateUserParam) (_ user.User, err error) {
	e, err := r.Get(id)
	if err != nil {
		return
	}

	if p.Name != nil {
		e.Name = *p.Name
	}
	if p.Email != nil {
		e.Email = *p.Email
	}
	if p.Password != nil {
		e.Password = *p.Password
	}
	if p.PostEventAvailabled != nil {
		e.PostEventAvailabled = *p.PostEventAvailabled
	}
	if p.Manage != nil {
		e.Manage = *p.Manage
	}
	if p.TwitterId.KeyExists() {
		if p.TwitterId.IsNull() {
			e.TwitterId = nil
		} else {
			e.TwitterId = *p.TwitterId.Value
		}
	}
	if p.GithubUsername.KeyExists() {
		if p.GithubUsername.IsNull() {
			e.GithubUsername = nil
		} else {
			e.GithubUsername = *p.GithubUsername.Value
		}
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
