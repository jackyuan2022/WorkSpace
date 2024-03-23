package dto

type GetDataListRequest[T any] struct {
	DataList   []T     `json:"data_list"`
	Pagination PageDTO `json:"page_info"`
}

type DataListResponse[T any] struct {
	DataList   []T     `json:"data_list"`
	Pagination PageDTO `json:"page_info"`
}

type DataRequest[T any] struct {
	Data T `json:"data"`
}

type DataResponse[T any] struct {
	Data T `json:"data"`
}

type PageDTO struct {
	PageSize    int  `json:"page_size"`
	PageNumber  int  `json:"page_number"`
	HasNextPage bool `json:"has_next_page"`
}
