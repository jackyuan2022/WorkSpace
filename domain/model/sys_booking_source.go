package model

import (
	core "github.com/jackyuan2022/workspace/core"
)

type BookingSource struct {
	core.DbBaseModel
	Name     string   `json:"name" gorm:"size:100;not null"`
	Category Category `json:"category" gorm:"foreignKey:category_id;references:id"`
}

func (entity *BookingSource) TableName() string {
	return "work_space_booking_source"
}
