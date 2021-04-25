package util

import (
	"math"
)

// Paginator struct for holding pagination info
type Paginator struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int64 `json:"total_page"`
	Offset      int64 `json:"offset"`
	Limit       int64 `json:"limit"`
	Page        int64 `json:"page"`
	PrevPage    int64 `json:"prev_page"`
	NextPage    int64 `json:"next_page"`
}

// PaginationData struct for returning pagination stat
type PaginationData struct {
	Total     int64 `json:"total" xml:"total"`
	Page      int64 `json:"page" xml:"page"`
	PerPage   int64 `json:"perPage" xml:"perPage"`
	Prev      int64 `json:"prev" xml:"prev"`
	Next      int64 `json:"next" xml:"next"`
	TotalPage int64 `json:"totalPage" xml:"totalPage"`
}

type PagedModel struct {
	Data     interface{}     `json:"data" xml:"data"`
	PageInfo *PaginationData `json:"pageInfo" xml:"pageInfo"`
}

func (p *Paginator) PagedData(data interface{}) *PagedModel {
	return &PagedModel{
		Data:     data,
		PageInfo: p.PaginationData(),
	}
}

// PaginationData returns PaginationData struct which
// holds information of all stats needed for pagination
func (p *Paginator) PaginationData() *PaginationData {
	data := PaginationData{
		Total:     p.TotalRecord,
		Page:      p.Page,
		PerPage:   p.Limit,
		Prev:      0,
		Next:      0,
		TotalPage: p.TotalPage,
	}
	if p.Page != p.PrevPage && p.TotalRecord > 0 {
		data.Prev = p.PrevPage
	}
	if p.Page != p.NextPage && p.TotalRecord > 0 && p.Page <= p.TotalPage {
		data.Next = p.NextPage
	}

	return &data
}

func Paging(page int64, limit int64, count int64) *Paginator {
	paginator := Paginator{}
	var offset int64

	if page < 1 {
		page = 1
	}
	if limit == 0 {
		limit = count
	}

	if page > 0 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}
	paginator.TotalRecord = count

	paginator.Page = page
	paginator.Offset = offset
	paginator.Limit = limit
	paginator.TotalPage = int64(math.Ceil(float64(count) / float64(limit)))
	if page > 1 {
		paginator.PrevPage = page - 1
	} else {
		paginator.PrevPage = page
	}
	if page == paginator.TotalPage {
		paginator.NextPage = page
	} else {
		paginator.NextPage = page + 1
	}

	return &paginator
}
