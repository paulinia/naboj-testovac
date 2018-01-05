package main

import (
	"fmt"
)

type CannotSolveError struct {
	id int
}

func (e CannotSolveError) Error() string {
	return fmt.Sprintf("User can't solve problem %v.", e.id)
}

type WrongPassword struct {
	user string
}

func (e WrongPassword) Error() string {
	return fmt.Sprintf("Wrong password")
}
