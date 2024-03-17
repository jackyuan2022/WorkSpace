package model

import (
	core "github.com/jackyuan2022/workspace/core"
)

type Category struct {
	core.DbBaseModel
	Name         string `json:"name" gorm:"size:100;not null"`
	CategoryType string `json:"category_type" gorm:"size:100"`
	Icon         string `json:"icon" gorm:"size:1000"`
	Order        int    `json:"order"`
}

func (entity *Category) TableName() string {
	return "work_space_category"
}
