package main

import "time"
import "fmt"

type Scoreboard []struct {
	Name   string
	Points int
	Last   time.Time
}

func (s Scoreboard) String() string {
	str := ""
	str += fmt.Sprintf("\tName\tPoints\tLast Submit\n")
	for i, u := range s {
		str += fmt.Sprintf("%v\t%v\t%v\t%v\n", i, u.Name, u.Points, u.Last)
	}
	return str
}
