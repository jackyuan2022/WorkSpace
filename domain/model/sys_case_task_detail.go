package model

import (
	"time"

	core "github.com/jackyuan2022/workspace/core"
)

type CaseTaskDetail struct {
	core.DbBaseModel
	CaseTaskId     string     `json:"case_task_id" gorm:"size:32"`
	Name           string     `json:"name" gorm:"size:100;not null"`
	CaseTime       time.Time  `json:"case_time" gorm:"not null"`
	ExpirationTime *time.Time `json:"expiration_time"`
	ExpirationDays int32      `json:"expiration_days"`
	Content        string     `json:"content" gorm:"size:1000"`
	UserId         string     `json:"user_id" gorm:"size:32"`
	CaseUser       User       `json:"case_user" gorm:"foreignKey:UserId;references:id"`
	Status         string     `json:"status" gorm:"size:100;default:ongoing"`
}

func (entity *CaseTaskDetail) TableName() string {
	return "work_space_case_task_detail"
}
