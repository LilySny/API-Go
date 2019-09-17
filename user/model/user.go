package model

import "time"

type User struct {
	ID           int       `json:id`
	Username     string    `json:username`
	Email        string    `json:email`
	RegisterDate time.Time `json:resgister-date`
	UpdateTime   time.Time `json:update-time`
}
