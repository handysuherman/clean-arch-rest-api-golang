package helper

import (
	"fmt"
	"math"
	"strconv"
)

const (
	defaultSize = 10
	defaultPage = 1
)

type PaginationResponse struct {
	TotalCount int64 `json:"totalCount"`
	TotalPages int64 `json:"totalPages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasMore    bool  `json:"hasMore"`
}

type PaginationTimeRangeResponse struct {
	TotalCount int64 `json:"totalCount"`
	TotalPages int64 `json:"totalPages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	From       int64 `json:"from"`
	To         int64 `json:"to"`
	HasMore    bool  `json:"hasMore"`
}

func (p *PaginationResponse) String() string {
	return fmt.Sprintf("TotalCount: %d, TotalPages: %d, Page: %d, Size: %d, HasMore: %v", p.TotalCount, p.TotalPages, p.Page, p.Size, p.HasMore)
}

func (p *PaginationTimeRangeResponse) String() string {
	return fmt.Sprintf("TotalCount: %d, TotalPages: %d, Page: %d, Size: %d, HasMore: %v, From: %v, To: %v", p.TotalCount, p.TotalPages, p.Page, p.Size, p.HasMore, p.From, p.To)
}

func NewPaginationResponse(totalCount int64, pq *Pagination) *PaginationResponse {
	return &PaginationResponse{
		TotalCount: totalCount,
		TotalPages: int64(pq.GetTotalPages(int(totalCount))),
		Page:       int64(pq.GetPage()),
		Size:       int64(pq.GetSize()),
		HasMore:    pq.GetHasMore(int(totalCount)),
	}
}

func NewPaginationTimeRangeResponse(from, to, totalCount int64, pq *Pagination) *PaginationTimeRangeResponse {
	return &PaginationTimeRangeResponse{
		TotalCount: totalCount,
		TotalPages: int64(pq.GetTotalPages(int(totalCount))),
		Page:       int64(pq.GetPage()),
		Size:       int64(pq.GetSize()),
		From:       from,
		To:         to,
		HasMore:    pq.GetHasMore(int(totalCount)),
	}
}

// Pagination query params
type Pagination struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

// NewPaginationQuery Pagination query constructor
func NewPaginationQuery(size int, page int) *Pagination {
	p := &Pagination{Size: defaultSize, Page: defaultPage}

	if size != 0 {
		p.Size = size
	}

	if page != 0 {
		p.Page = page
	}

	return p
}

func NewPaginationFromQueryParams(size string, page string) *Pagination {
	p := &Pagination{Size: defaultSize, Page: 1}

	if sizeNum, err := strconv.Atoi(size); err == nil && sizeNum != 0 {
		p.Page = sizeNum
	}

	if pageNum, err := strconv.Atoi(page); err == nil && pageNum != 0 {
		p.Page = pageNum
	}

	return p
}

// SetSize Set page size
func (q *Pagination) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	q.Size = n

	return nil
}

// SetPage Set page number
func (q *Pagination) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

// SetOrderBy Set order by
func (q *Pagination) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

// GetOffset Get offset
func (q *Pagination) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

// GetLimit Get limit
func (q *Pagination) GetLimit() int {
	return q.Size
}

// GetOrderBy Get OrderBy
func (q *Pagination) GetOrderBy() string {
	return q.OrderBy
}

// GetPage Get OrderBy
func (q *Pagination) GetPage() int {
	return q.Page
}

// GetSize Get OrderBy
func (q *Pagination) GetSize() int {
	return q.Size
}

// GetQueryString get query string
func (q *Pagination) GetQueryString() string {
	if q.OrderBy == "" {
		return fmt.Sprintf("page=%d&size=%d", q.GetPage(), q.GetSize())
	}
	return fmt.Sprintf("page=%d&size=%d&orderBy=%s", q.GetPage(), q.GetSize(), q.GetOrderBy())
}

// GetTotalPages Get total pages int
func (q *Pagination) GetTotalPages(totalCount int) int {
	d := float64(totalCount) / float64(q.GetSize())
	return int(math.Ceil(d))
}

// GetHasMore Get has more
func (q *Pagination) GetHasMore(totalCount int) bool {
	return float64(q.GetPage()) < float64(float64(totalCount)/float64(q.GetSize()))
}
