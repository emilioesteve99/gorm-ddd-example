package commonDomainModels

type PaginatedItemsMeta struct {
	CurrentPage  int `json:"currentPage"`
	ItemCount    int `json:"itemCount"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalItems   int `json:"totalItems"`
	TotalPages   int `json:"totalPages"`
}

type PaginatedItems[TItem any] struct {
	Items []TItem            `json:"items"`
	Meta  PaginatedItemsMeta `json:"meta"`
}

type PaginationContext struct {
	Limit      int
	Page       int
	TotalItems int
}
