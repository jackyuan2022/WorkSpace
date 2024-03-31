package model

import (
	"time"

	core "github.com/jackyuan2022/workspace/core"
)

type CaseTask struct {
	core.DbBaseModel
	Name            string           `json:"name" gorm:"size:100;not null"`
	CaseTime        time.Time        `json:"case_time" gorm:"not null"`
	Content         string           `json:"content" gorm:"size:1000"`
	UserId          string           `json:"user_id" gorm:"size:32"`
	CaseUser        User             `json:"case_user" gorm:"foreignKey:UserId;references:Id"`
	Status          string           `json:"status" gorm:"size:100;default:ongoing"`
	CaseTaskDetails []CaseTaskDetail `json:"case_task_details" gorm:"foreignKey:CaseTaskId;references:Id"`
}

func (entity *CaseTask) TableName() string {
	return "work_space_case_task"
}
