package errors

import "fmt"

type DBError interface {
	CustomError()
	Error() string
}

type NotFound struct {
	DBError
	ID string
}

func (n NotFound) Error() string {
	return fmt.Sprint("Not found id: " + n.ID)
}
