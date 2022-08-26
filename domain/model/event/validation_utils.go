package event

import (
	"errors"
)

// Errors
var (
	ErrEventTitleCannotBeEmpty   = errors.New("Event title cannot be empty")
	ErrDocumentNameCannotBeEmpty = errors.New("Event document name cannot be empty")
	ErrUrlCannotBeEmpty          = errors.New("Event document url cannot be empty")
)

func validateTitle(title string) error {
	// 空文字チェック
	if title == "" {
		return ErrEventTitleCannotBeEmpty
	}
	return nil
}

func validateDocumentName(name string) error {
	// 空文字チェック
	if name == "" {
		return ErrDocumentNameCannotBeEmpty
	}
	return nil
}

func validateUrl(url string) error {
	// 空文字チェック
	if url == "" {
		return ErrUrlCannotBeEmpty
	}
	return nil
}

func validateEventId(repo EventRepository, id string) error {
	_, err := repo.Get(id)
	return err
}
