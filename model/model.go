package model

import "time"

type User struct {
	UserID       string     `son:"_id,omitempty"`
	Name         string     `bson:"name,omitempty"`
	PasswordHash string     `bson:"password,omitempty"`
	LastLogin    *time.Time `bson:"lastlogtm,omitempty"`
}
