package model

import (
	"database/sql"
	"time"

	core "github.com/jackyuan2022/workspace/core"
)

type Booking struct {
	core.DbBaseModel
	Title            string         `json:"title" gorm:"size:100;not null"`
	BookingStartTime time.Time      `json:"booking_start_time" gorm:"not null"`
	BookingEndTime   sql.NullTime   `json:"booking_end_time"`
	Content          sql.NullString `json:"content" gorm:"size:1000"`
	BookingSourceId  string         `json:"booking_source_id" gorm:"size:32"`
	BookingSource    BookingSource  `json:"booking_source" gorm:"foreignKey:BookingSourceId;references:id"`
	UserId           string         `json:"user_id" gorm:"size:32"`
	BookingUser      User           `json:"user" gorm:"foreignKey:UserId;references:id"`
	Status           string         `json:"status" gorm:"size:100;default:ongoing"`
}

func (entity *Booking) TableName() string {
	return "work_space_booking"
}
