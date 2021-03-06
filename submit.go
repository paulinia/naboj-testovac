package main

import "time"
import "fmt"

type Submit struct {
	t      time.Time
	task   int
	points int
}

func (c *Contest) pointValue(task int, sumbits []Submit) int {
	fmt.Println(sumbits)
	cnt := 0
	for _, s := range sumbits {
		fmt.Printf("%v ", s.task)
		if s.task == task {
			cnt++
			if s.points > 0 {
				return 0
			}
		}
	}
	fmt.Println("Pokusov bolo: ", cnt)
	if cnt > len(c.scoring) {
		return c.scoring[len(c.scoring)-1]
	}
	return c.scoring[cnt]
}
