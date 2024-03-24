package dto

import (
	"time"
)

type GetBookingListRequest struct {
	BookingSourceId string  `json:"booking_source_id"`
	UserId          string  `json:"user_id"`
	Pagination      PageDTO `json:"page_info"`
}

type BookingDTO struct {
	Id               string           `json:"id"`
	Title            string           `json:"title"`
	Content          string           `json:"content"`
	UserId           string           `json:"user_id"`
	BookingUser      UserDTO          `json:"booking_user"`
	BookingSourceId  string           `json:"booking_source_id"`
	BookingSource    BookingSourceDTO `json:"booking_source"`
	Status           string           `json:"status"`
	BookingStartTime time.Time        `json:"booking_start_time" time_format:"2006-01-02 12:00:00" time_utc:"true"`
	BookingEndTime   *time.Time       `json:"booking_end_time" time_format:"2006-01-02 12:00:00" time_utc:"true"`
}
