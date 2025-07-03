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

func NewPagerInput(pageNumber int, pageSize int) *PagerInput {
	return &PagerInput{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
}
