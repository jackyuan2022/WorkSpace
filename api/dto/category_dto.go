package dto

type GetCategoryListRequest struct {
	CategoryType string  `json:"category_type"`
	Pagination   PageDTO `json:"page_info"`
}

type CategoryDTO struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Order        int    `json:"order"`
	CategoryType string `json:"category_type"`
}
