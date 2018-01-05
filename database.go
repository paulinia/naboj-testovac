package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type Database struct {
	n        int
	begin    int
	end      int
	problems []Problem
}

func (d Database) String() string {
	pole := make([]string, 0)
	for _, s := range d.problems {
		pole = append(pole, s.String())
	}
	return fmt.Sprintf("Database of %v problems.\n[%v; %v]\nProblems are: %v", d.n, d.begin, d.end, strings.Join(pole, "\n"))
}

func (d *Database) getProblems(n int, typ []string) []Problem {
	cnt := 0
	selected := make([]Problem, 0)
	possible := make([]int, 0)
	for i, P := range d.problems {
		ok := false
		for _, s := range typ {
			if s == P.typ {
				ok = true
				break
			}
		}
		if !ok {
			continue
		}
		possible = append(possible, i)
		cnt++

	}

	id, last := 0, 0
	for i := 0; i < n; i++ {
		upTo := id
		for j, idd := range possible[id:] {
			if d.problems[idd].level-last > 4 {
				upTo = j + id
			}
		}
		if cnt-id >= n-i {
			for _, idd := range possible[id:] {
				selected = append(selected, d.problems[idd])
			}
			return selected
		}
		count := upTo - id
		if cnt-count < n-i {
			upTo = cnt - n - i - 1
			count = upTo - id
		}
		index := rand.Intn(count)
		selected = append(selected, d.problems[possible[index]])
	}

	return selected

}

func (d *Database) addProblem(p Problem) {
	d.n++
	d.problems = append(d.problems, p)
	if d.begin > p.level || d.begin == 0 {
		d.begin = p.level
	}
	if d.end < p.level {
		d.end = p.level
	}

	sort.Slice(d.problems, func(i, j int) bool {
		if d.problems[i].level < d.problems[j].level {
			return true
		}
		return false
	})

}
