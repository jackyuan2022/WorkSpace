package dto

type GetCategoryListResponse struct {
	CategoryList []CategoryDTO `json:"categories"`
	HasNextPage  bool          `json:"has_next_page"`
}

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

type CategoryResponse struct {
	Category CategoryDTO `json:"category"`
}

type CategoryRequest struct {
	Category CategoryDTO `json:"category"`
}
