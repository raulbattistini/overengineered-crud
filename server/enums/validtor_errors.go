package enums

type ValidtorErrorMessage string

const (
	ValidtorNoPostRecived   ValidtorErrorMessage = "no post body was recived"
	ValidtorNoIdRecived     ValidtorErrorMessage = "no id was recived"
	ValidtorInvalidIdFormat ValidtorErrorMessage = "invalid id format"
	ValidtorEmptyTitle      ValidtorErrorMessage = "no title was recived"
	ValidtorInvalidTitle    ValidtorErrorMessage = "invalid title format"
	ValidtorEmptyContent    ValidtorErrorMessage = "invalid content format"
)
