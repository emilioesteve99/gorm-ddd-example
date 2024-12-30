package appErrors

type ErrorMessage string

const (
	UnknownMsg             ErrorMessage = "Unknown error"
	ForeignKeyViolationMsg ErrorMessage = "Foreign key violation"
	DuplicatedEntityMsg    ErrorMessage = "Duplicated entity"
)
