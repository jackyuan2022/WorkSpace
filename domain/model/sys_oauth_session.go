package model

import (
	core "github.com/jackyuan2022/workspace/core"
)

type OAuthSession struct {
	core.DbBaseModel
	UserId       string `json:"user_id" gorm:"size:32"`
	Mobile       string `json:"mobile" gorm:"size:100"`
	UserName     string `json:"user_name"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	ClientIp     string `json:"client_ip"`
	IsBlocked    bool   `json:"is_blocked"`
	ExpiredAt    int64  `json:"expired_at"`
}

func (entity *OAuthSession) TableName() string {
	return "work_space_oauth_session"
}
