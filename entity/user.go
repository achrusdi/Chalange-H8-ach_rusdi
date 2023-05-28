package entity

import "time"

type User struct {
	ID         int
	Username   string
	Password   string
	Created_at time.Time
}
