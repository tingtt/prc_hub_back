package event

import (
	"prc_hub_back/domain/model/user"
	"time"
)

type Event struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	User        *user.User       `json:"user"`
	Description *string          `json:"description"`
	Location    *string          `json:"location,omitempty"`
	Datetimes   []EventDatetime  `json:"datetimes"`
	Documents   *[]EventDocument `json:"documents,omitempty"`
	Published   bool             `json:"published"`
	Completed   bool             `json:"completed"`
	UserId      string           `json:"user_id"`
}

type EventDatetime struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type EventDocument struct {
	EventId string `json:"event_id"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}
