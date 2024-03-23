package dto

import (
	"time"
)

type GetBookingListRequest struct {
	CategoryId string `json:"category_id"`
	UserId     string `json:"user_id"`
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
}

type BookingDTO struct {
	Id               string      `json:"id"`
	Title            string      `json:"title"`
	Content          string      `json:"content"`
	UserId           string      `json:"user_id"`
	BookingUser      UserDTO     `json:"booking_user"`
	CategoryId       string      `json:"category_id"`
	Category         CategoryDTO `json:"category"`
	Status           string      `json:"status"`
	BookingStartTime time.Time   `json:"booking_start_time"`
	BookingEndTime   *time.Time  `json:"booking_end_time"`
}
