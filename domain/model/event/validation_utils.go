package event

import (
	"errors"
)

// Errors
var (
	ErrValidateEventTitleCannotBeEmpty   = errors.New("Event title cannot be empty")
	ErrValidateDocumentNameCannotBeEmpty = errors.New("Event document name cannot be empty")
	ErrValidateUrlCannotBeEmpty          = errors.New("Event document url cannot be empty")
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
	_, err := repo.Get(id)
	return err
}
