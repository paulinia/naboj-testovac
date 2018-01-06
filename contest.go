package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"time"
)

type Contest struct {
	n          int
	beginP     int
	started    bool
	start      time.Time
	problemset []Problem
	users      map[string]User
	scoring    []int
}

func (c Contest) String() string {
	return fmt.Sprintf("Contest of %v / %v problems [%v] started at %v.\n"+
		"Problem set: %v\nUsers:%v\n"+
		"And scoring si %v\n", c.n, c.beginP, c.started, c.start, c.problemset, c.users, c.scoring)
}

func validContest(p []Problem) bool {
	last := 0
	for _, pr := range p {
		fmt.Printf("%v ", pr.level)
		if pr.level-last > 4 {
			return false
		}
		last = pr.level
	}
	return true
}

func generateContest(n, bp int, typ []string, scoring []int, D *Database) Contest {
	dp := make([]Problem, 0)
	for {
		dp = D.getProblems(n, typ)
		fmt.Println("Moze byt")
		if validContest(dp) {
			break
		}
	}
	return Contest{
		n,
		bp,
		false,
		time.Now(),
		dp,
		make(map[string]User),
		scoring,
	}
}

func (c *Contest) addUser(name string, password string) {
	newU := User{
		name,
		sha256.Sum224([]byte(password)),
		0,
		make([]int, c.beginP),
		make([]Submit, 0),
	}
	for i := 0; i < c.beginP; i++ {
		newU.avialable[i] = i
	}
	c.users[name] = newU
}

func (c *Contest) begin() {
	c.started = true
	c.start = time.Now()
}

func (c *Contest) end() {
	c.started = false
}

func (c *Contest) getScoreboard() Scoreboard {
	scoreboard := make(Scoreboard, len(c.users))
	users := make([]string, 0)

	for name := range c.users {
		users = append(users, name)
	}

	sort.Slice(users, func(i, j int) bool {
		return (c.users[users[i]].points > c.users[users[j]].points)
	})

	for i, id := range users {
		scoreboard[i].name = c.users[id].name
		if len(c.users[id].submits) > 0 {
			scoreboard[i].last = c.users[id].submits[len(c.users[id].submits)-1].t
		}
		scoreboard[i].points = c.users[id].points
	}
	return scoreboard
}

func (c *Contest) show(user, password string, id int) (string, error) {
	if sha256.Sum224([]byte(password)) != c.users[user].password {
		return "", WrongPassword{user}
	}
	for _, u := range c.users {
		if u.name != user {
			continue
		}
		for _, p := range u.avialable {
			if p == id {
				return c.problemset[id].statement, nil
			}
		}
	}
	return "", DontHaveAccesError{user, id}
}

func (c *Contest) submit(user, password string, id int, sol string) (points int, er error) {
	if sha256.Sum224([]byte(password)) != c.users[user].password {
		points = 0
		er = WrongPassword{user}
		return
	}
	new := make([]int, 0)
	u := c.users[user]
	er = CannotSolveError{id}

	for _, p := range c.users[user].avialable {
		if p != id {
			new = append(new, p)
			continue
		}
		if c.problemset[p].solved(sol) {
			points = c.pointValue(id, c.users[user].submits)
			fnc := func() {
				u.avialable = append(new, new[len(new)-1]+1)
				c.users[user] = u
			}
			defer fnc()
			u.points += points
			u.submits = append(u.submits, Submit{time.Now(), id, points})
		}
		er = nil
	}
	return
}
