package main

import "time"
import "fmt"

type Scoreboard []struct {
	name   string
	points int
	last   time.Time
}

func (s Scoreboard) String() string {
	str := ""
	str += fmt.Sprintf("\tName\tPoints\tLast Submit\n")
	for i, u := range s {
		str += fmt.Sprintf("%v\t%v\t%v\t%v\n", i, u.name, u.points, u.last)
	}
	return str
}
