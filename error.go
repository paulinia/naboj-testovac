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

type DontHaveAccesError struct {
	user string
	id   int
}

func (e DontHaveAccesError) Error() string {
	return fmt.Sprintf("%v doesn't solve problem %v.", e.user, e.id)
}

type WrongPassword struct {
	user string
}

func (e WrongPassword) Error() string {
	return fmt.Sprintf("Wrong password")
}

type WrongAnswer struct {
}

func (e WrongAnswer) Error() string {
	return "Wrong answer."
}
