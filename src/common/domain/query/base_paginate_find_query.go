package commonDomainQueries

type PaginationOptions struct {
	Limit int
	Page  int
}

type BasePaginateFindQuery struct {
	PaginationOptions PaginationOptions
}

func (q BasePaginateFindQuery) Limit() int {
	return q.PaginationOptions.Limit
}

func (q BasePaginateFindQuery) Page() int {
	return q.PaginationOptions.Page
}
