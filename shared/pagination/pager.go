package paginationShared

type PagerInput struct {
	PageNumber int `json:"page" form:"page"`
	PageSize   int `json:"size" form:"size"`
}

func (p *PagerInput) IsValid() bool {
	return p.PageNumber > 0 && p.PageSize > 0
}

type PagerOutput struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

func (o SortOrder) IsValid() bool {
	return o == SortOrderAsc || o == SortOrderDesc
}

func NewPagerInput(pageNumber int, pageSize int) *PagerInput {
	return &PagerInput{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
}
