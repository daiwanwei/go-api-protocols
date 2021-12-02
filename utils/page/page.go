package page

import "errors"

var ErrCovertContent = errors.New("fail to covert content")

type Page struct {
	Size    int         `json:"size"`
	Page    int         `json:"page"`
	Total   int64       `json:"total"`
	Content interface{} `json:"content"`
}

type Pageable struct {
	Size int            `json:"size"`
	Page int            `json:"page"`
	Sort map[string]int `json:"sorts"`
}
