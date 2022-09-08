package webhook

import (
	"fmt"
	"prc_hub_back/domain/model/event"
	"strings"
)

func NotifyEvent(e event.Event) error {
	// メッセージを生成
	msg := "イベント情報\n\n%s: %s\n\n"
	msgParams := []interface{}{"勉強会", e.Name}
	if e.Description != nil {
		msg += "%s\n\n"
		msgParams = append(msgParams, *e.Description)
	}
	msg += "%s/events/%d"
	msgParams = append(msgParams, frontUrl, e.Id)

	return NotifyToAllProviders(fmt.Sprintf(msg, msgParams...))
}

func NotifyEventDocuments(e event.Event) error {
	if e.Documents != nil && len(*e.Documents) != 0 {
		// メッセージを生成
		msg := "イベント資料\n\n%s: %s\n\n"
		msgParams := []interface{}{"勉強会", e.Name}

		for _, d := range *e.Documents {
			msg += "%s\n%s\n\n"
			msgParams = append(msgParams, d.Name, d.Url)
		}
		msg = strings.TrimSuffix(msg, "\n\n")

		return NotifyToAllProviders(fmt.Sprintf(msg, msgParams...))
	}
	return nil
}
