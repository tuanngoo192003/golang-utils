package utils

type PaginationResponse[T any] struct {
    Page        int   `json:"page"`
    PerPage     int   `json:"perPage"`
    Data        []T   `json:"data"`
    TotalRecord int64 `json:"totalRecord"`
    TotalPage   int64 `json:"totalPage"`
}
