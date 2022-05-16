package models

import "time"

type userType uint8

const (
	user  userType = 1
	robot userType = 2
)

type Token struct {
	ID uint `gorm:"primarykey"`

	// grant client info
	ClientID    string
	RedirectURI string
	State       string

	// token basic info
	// Code authorize_code/access_token/refresh-token
	Code      string
	CreatedAt time.Time
	ExpiresIn time.Duration
	Scope     string

	UsrOrRobotIdentity string
}

type OauthServerInfo struct {
	ID          uint
	ClientID    string
	RedirectURI string
	HomeURL     string
}

type ClientSecret struct {
	ID           uint
	ClientID     string
	RedirectURL  string
	ClientSecret string
	CreateTime   time.Time
	CreateBy     uint
}
