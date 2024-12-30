package commonDomainQueries

type PaginateFindQuery interface {
	Limit() int
	Page() int
}
