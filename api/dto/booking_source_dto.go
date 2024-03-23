package dto

type GetBookingSourceListRequest struct {
	CategoryId string  `json:"category_id"`
	Pagination PageDTO `json:"page_info"`
}

type BookingSourceDTO struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Category CategoryDTO `json:"category"`
}
