package paginationShared

type PagerInput struct {
	PageNumber int
	PageSize   int
}

type PagerOutput struct {
	TotalItems  int
	TotalPages  int
	CurrentPage int
	PageSize    int
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
