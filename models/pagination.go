package models

import ()

const (
	PAGINATION_DEFAULT_LENGTH int = 10
)

type Pagination struct {
	StartCount int
	EndCount   int
	TotalCount int64
	Page       int
	MaxPage    int
	Pages      []int
	Length     int
	Range      int
}

func NewPagination(page int, length int, totalCount int64) *Pagination {
	p := Pagination{}

	p.Length = length
	p.Range = 3
	p.MaxPage = int(totalCount / int64(length))
	if int64(p.MaxPage*length) < totalCount {
		p.MaxPage += 1
	}

	p.Page = page
	if p.Page <= 0 || p.Page > p.MaxPage {
		p.Page = 1
	}

	p.StartCount = ((p.Page - 1) * length) + 1
	p.TotalCount = totalCount
	if p.TotalCount <= 0 {
		p.StartCount = 0
	}

	startPage := p.Page - p.Range
	if startPage <= 0 {
		startPage = 1
	}

	endPage := p.Page + p.Range
	if endPage > p.MaxPage {
		endPage = p.MaxPage
	}

	pages := make([]int, 0)
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, i)
	}
	p.Pages = pages

	return &p
}
