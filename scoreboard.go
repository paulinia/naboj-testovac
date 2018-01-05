package main

import "time"
import "fmt"

type Scoreboard []struct {
	name   string
	points int
	last   time.Time
}

func (s Scoreboard) show() {
	fmt.Printf("\tName\tPoints\tLast Submit\n")
	for i, u := range s {
		fmt.Printf("%v\t%v\t%v\t%v\n", i, u.name, u.points, u.last)
	}
}
