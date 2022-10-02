package event

import (
	"prc_hub_back/domain/model/user"
	"time"
)

type Event struct {
	Id          string          `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Description *string         `json:"description,omitempty" db:"description"`
	Location    *string         `json:"location,omitempty" db:"location"`
	Datetimes   []EventDatetime `json:"datetimes" db:"datetimes"`
	Published   bool            `json:"published" db:"published"`
	Completed   bool            `json:"completed" db:"completed"`
	UserId      string          `json:"user_id" db:"user_id"`
}

type EventDatetime struct {
	Start time.Time `json:"start" db:"start"`
	End   time.Time `json:"end" dh:"end"`
}

type EventDocument struct {
	EventId string `json:"event_id" db:"event_id"`
	Id      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Url     string `json:"url" db:"url"`
}

type EventEmbed struct {
	Event
	User      *user.User       `json:"user,omitempty"`
	Documents *[]EventDocument `json:"documents,omitempty"`
}
