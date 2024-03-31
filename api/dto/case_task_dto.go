package dto

type GetCaseTaskListRequest struct {
	UserId     string  `json:"user_id"`
	CategoryId string  `json:"category_id"`
	Pagination PageDTO `json:"page_info"`
}

type CaseTaskDTO struct {
	Id              string              `json:"id"`
	Name            string              `json:"name"`
	CaseTime        int64               `json:"case_time"`
	Content         string              `json:"content"`
	CategoryId      string              `json:"category_id"`
	UserId          string              `json:"user_id"`
	CaseUser        UserDTO             `json:"case_user"`
	Status          string              `json:"status"`
	CaseTaskDetails []CaseTaskDetailDTO `json:"case_task_details"`
}

type CaseTaskDetailDTO struct {
	Id             string  `json:"id"`
	CaseTaskId     string  `json:"case_task_id"`
	Name           string  `json:"name"`
	CaseTime       int64   `json:"case_time"`
	ExpirationTime *int64  `json:"expiration_time"`
	ExpirationDays int32   `json:"expiration_days"`
	Content        string  `json:"content"`
	UserId         string  `json:"user_id"`
	CaseUser       UserDTO `json:"case_user"`
	Status         string  `json:"status"`
}
