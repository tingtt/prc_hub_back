package user_mysql

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type (
	Repository struct{}
)

// Singleton field
var dsn string

func InitRepository(user string, password string, host string, port uint, db string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, db)
}

// MySQLサーバーに接続
func Open() (*sqlx.DB, error) {
	if dsn == "" {
		return nil, errors.New("dsn does not set")
	}
	return sqlx.Open("mysql", dsn)
}