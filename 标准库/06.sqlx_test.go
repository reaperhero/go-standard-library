package stand

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

type Project struct {
	ID     int
	UserID int `db:"user_id"`
	Name   string
}

var (
	conn *sqlx.DB
)

// slelect any
func Test_SelectAny(t *testing.T) {
	sql := "select id,user_id,name from projects where user_id=?"
	projects := []Project{}
	conn.Select(&projects, sql, 1)
}


// update
