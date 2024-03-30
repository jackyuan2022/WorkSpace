package model

import (
	"time"

	core "github.com/jackyuan2022/workspace/core"
)

type CaseTask struct {
	core.DbBaseModel
	Title            string        `json:"title" gorm:"size:100;not null"`
	BookingStartTime time.Time     `json:"booking_start_time" gorm:"not null"`
	BookingEndTime   *time.Time    `json:"booking_end_time"`
	Content          string        `json:"content" gorm:"size:1000"`
	BookingSourceId  string        `json:"booking_source_id" gorm:"size:32"`
	BookingSource    BookingSource `json:"booking_source" gorm:"foreignKey:BookingSourceId;references:id"`
	UserId           string        `json:"user_id" gorm:"size:32"`
	BookingUser      User          `json:"user" gorm:"foreignKey:UserId;references:id"`
	Status           string        `json:"status" gorm:"size:100;default:ongoing"`
}

func (entity *CaseTask) TableName() string {
	return "work_space_case_task"
}
