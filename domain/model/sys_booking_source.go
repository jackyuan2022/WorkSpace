package model

import (
	core "github.com/jackyuan2022/workspace/core"
)

type BookingSource struct {
	core.DbBaseModel
	Name       string   `json:"name" gorm:"size:100;not null"`
	CategoryId string   `json:"category_id" gorm:"size:32"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId;references:id"`
}

func (entity *BookingSource) TableName() string {
	return "work_space_booking_source"
}
