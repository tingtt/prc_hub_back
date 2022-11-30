package event

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"prc_hub_back/domain/model/user"
	"prc_hub_back/domain/model/util"
	"time"
)

// Errors
var (
	ErrCannotCreateEvent = errors.New("sorry, you cannot create `event`")
)

type CreateEventParam struct {
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	Datetimes   []CreateEventDatetimeParam `json:"datetimes"`
	Location    *string                    `json:"location,omitempty"`
	Published   bool                       `json:"published"`
	Completed   bool                       `json:"completed"`
}

type CreateEventDatetimeParam struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (p CreateEventDatetimeParam) validate() error {
	// フィールドの検証
	err := validateEventDatetime(time.Time(p.Start), time.Time(p.End))
	if err != nil {
		return err
	}
	return nil
}

func (p CreateEventParam) validate(requestUser user.User) error {
	/**
	 * フィールドの検証
	**/
	// `Name`
	err := validateTitle(p.Name)
	if err != nil {
		return err
	}
	// `Datetimes`
	if len(p.Datetimes) == 0 {
		return ErrValidateDocumentNameCannotBeEmpty
	}
	for _, d := range p.Datetimes {
		err = d.validate()
		if err != nil {
			return err
		}
	}

	// 権限の検証
	if !requestUser.Admin && !requestUser.Manage && !requestUser.PostEventAvailabled {
		// `Admin`・`Manage`・`PostEventAvailabled`のいずれでもない場合は`Event`作成不可
		return ErrCannotCreateEvent
	}

	return nil
}

func CreateEvent(p CreateEventParam, requestUser user.User) (Event, error) {
	// バリデーション
	err := p.validate(requestUser)
	if err != nil {
		return Event{}, err
	}

	var datetimes []EventDatetime
	for _, d := range p.Datetimes {
		datetimes = append(datetimes, EventDatetime{
			Start: d.Start.UTC(),
			End:   d.End.UTC(),
		})
	}

	// TODO: UUID -> LastInsertedId()
	e := Event{
		Id:          util.UUID(),
		Name:        p.Name,
		Description: p.Description,
		Location:    p.Location,
		Datetimes:   datetimes,
		Published:   p.Published,
		Completed:   p.Completed,
		UserId:      requestUser.Id,
	}

	// MySQLサーバーに接続
	d, err := OpenMysql()
	if err != nil {
		return Event{}, err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer d.Close()

	// トランザクション開始
	tx, err := d.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return Event{}, err
	}
	defer func() {
		// return時にトランザクションの後処理
		//* 102行目の`defer`より先に実行される
		if err != nil {
			// 失敗時はロールバック
			tx.Rollback()
		} else {
			// 成功時はコミット
			tx.Commit()
		}
	}()

	// `events`テーブルに追加
	_, err = tx.NamedExec(
		`INSERT INTO events
			(id, name, description, location, published, completed, user_id)
		VALUES
			(:id, :name, :description, :location, :published, :completed, :user_id)`,
		e,
	)
	if err != nil {
		return Event{}, err
	}

	// `event_datetimes`テーブルに追加
	// TODO: for文にする
	_, err = tx.NamedExec(
		fmt.Sprintf(
			`INSERT INTO event_datetimes
				(event_id, start, end)
			VALUES
				("%s", :start, :end)`,
			e.Id,
		),
		e.Datetimes,
	)
	if err != nil {
		return Event{}, err
	}

	return e, nil
}
