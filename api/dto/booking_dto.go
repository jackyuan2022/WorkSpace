package dto

import (
	"time"
)

type GetBookingListResponse struct {
	BookingList []BookingDTO `json:"bookings"`
	HasNextPage bool         `json:"has_next_page"`
}

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
	BookingUser      UserDTO     `json:"booking_user"`
	Category         CategoryDTO `json:"category"`
	Status           string      `json:"status"`
	BookingStartTime time.Time   `json:"booking_start_time"`
	BookingEndTime   *time.Time  `json:"booking_end_time"`
}

type BookingResponse struct {
	Booking BookingDTO `json:"booking"`
}

type BookingRequest struct {
	Booking BookingDTO `json:"booking"`
}
