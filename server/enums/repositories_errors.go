package enums

type RepositoryErrorMessage string

const (
	RepositoryConnectionNotFound RepositoryErrorMessage = "missing database connection"
	RepositoryRowNotFound        RepositoryErrorMessage = "no records founded"
)
