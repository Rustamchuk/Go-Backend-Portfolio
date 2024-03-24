package model

import "time"

type RequestStat struct {
	UserId int64     `db:"user_id"`
	Time   time.Time `db:"time"`
}
