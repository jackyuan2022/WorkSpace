package dto

type GetCategoryListRequest struct {
	CategoryType string `json:"category_type"`
	PageSize     int    `json:"page_size"`
	PageNumber   int    `json:"page_number"`
}

type CategoryDTO struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	Order        int    `json:"order"`
	CategoryType string `json:"category_type"`
}
