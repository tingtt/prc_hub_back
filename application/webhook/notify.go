package webhook

import (
	"fmt"
	"prc_hub_back/domain/model/event"
	"strings"
)

func NotifyEvent(e event.EventEmbed) error {
	// メッセージを生成
	msg := "イベント情報\n\n%s: %s\n\n"
	msgParams := []interface{}{"勉強会", e.Name}
	if e.Description != nil {
		for _, s := range strings.Split(*e.Description, "\\n") {
			msg += "%s\n"
			msgParams = append(msgParams, s)
		}
		msg += "\n"
	}
	msg += "%s/events/%s"
	msgParams = append(msgParams, strings.TrimSuffix(frontUrl, "/"), e.Id)

	return NotifyToAllProviders(fmt.Sprintf(msg, msgParams...))
}

func NotifyEventDocuments(e event.EventEmbed) error {
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
