package event

import (
	"errors"
	"time"
)

// Errors
var (
	ErrValidateEventTitleCannotBeEmpty         = errors.New("Event title cannot be empty")
	ErrValidateDocumentNameCannotBeEmpty       = errors.New("Event document name cannot be empty")
	ErrValidateUrlCannotBeEmpty                = errors.New("Event document url cannot be empty")
	ErrValidateEventDatetimesCannotBeEmpty     = errors.New("Event datetime cannot be empty")
	ErrValidateEventDatetimeStartMustBeforeEnd = errors.New("Event start datetime must be before end datetime")
)

func validateTitle(title string) error {
	// 空文字チェック
	if title == "" {
		return ErrValidateEventTitleCannotBeEmpty
	}
	return nil
}

func validateDocumentName(name string) error {
	// 空文字チェック
	if name == "" {
		return ErrValidateDocumentNameCannotBeEmpty
	}
	return nil
}

func validateUrl(url string) error {
	// 空文字チェック
	if url == "" {
		return ErrValidateUrlCannotBeEmpty
	}
	return nil
}

func validateEventId(repo EventRepository, id string) error {
	_, err := repo.Get(id, GetEventQueryParam{})
	return err
}

func validateEventDatetime(start time.Time, end time.Time) error {
	if !start.Before(end) {
		return ErrValidateEventDatetimeStartMustBeforeEnd
	}
	return nil
}
