package user_inmemory

import (
	"prc_hub_back/domain/model/user"
	"strings"
)

func (r Repository) GetList(q user.GetUserListQueryParam) (_ []user.User, err error) {
	res := []user.User{}
	for _, u := range data {
		flag := true
		if q.Name != nil && *q.Name != u.Name {
			flag = false
		}
		if q.NameContain != nil && !strings.Contains(u.Name, *q.NameContain) {
			flag = false
		}
		if q.PostEventAvailabled != nil && *q.PostEventAvailabled != u.PostEventAvailabled {
			flag = false
		}
		if q.Manage != nil && *q.Manage != u.Manage {
			flag = false
		}
		if q.Admin != nil && *q.Admin != u.Admin {
			flag = false
		}

		if flag {
			res = append(res, u)
		}
	}
	return res, nil
}
