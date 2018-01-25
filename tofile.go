package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func (p *Problem) write(w *bufio.Writer) {
	fmt.Fprintln(w, p.statement)
	fmt.Fprintln(w, "ENDS")
	fmt.Fprintln(w, p.result)
	fmt.Fprintln(w, p.level)
	fmt.Fprintln(w, p.typ)
	fmt.Fprintln(w, p.imag)
}

func (d *Database) writeToFile(w *bufio.Writer) {
	for _, p := range d.problems {
		(&p).write(w)
	}
	fmt.Fprintln(w, "END")
	w.Flush()
}

func (d *Database) readDatabase(r *bufio.Reader) {
	for {
		s := getLine(r)
		if s == "END" {
			return
		}

		statement := ""

		for {
			if s == "ENDS" {
				break
			}
			statement += s
			s = getLine(r)
		}
		res := toResult(getLine(r))
		var lev int
		fmt.Sscan(getLine(r), &lev)
		// typ := getLine(r)
		d.addProblem(Problem{
			statement,
			res,
			lev,
			getLine(r),
			getLine(r),
		})
	}
}

func (c *Contest) read(r *bufio.Reader) {
	fmt.Sscan(getLine(r), &(c.start))
	fmt.Sscan(getLine(r), &(c.beginP))
	c.users = make(map[string]User)
	for {
		s := getLine(r)
		if s == "ENDP" {
			break
		}
		statement := ""
		for {
			if s == "ENDS" {
				break
			}
			statement += s
			s = getLine(r)
		}
		res := toResult(getLine(r))
		var lev int
		fmt.Sscan(getLine(r), &lev)
		c.problemset = append(c.problemset, Problem{
			statement,
			res,
			lev,
			getLine(r),
			getLine(r),
		})
	}
	for {
		s := getLine(r)
		if s == "ENDU" {
			break
		}
		name := s
		password := getLine(r)
		var points int
		fmt.Sscan(getLine(r), &points)
		reader := strings.NewReader(getLine(r))
		avialable := make([]int, c.beginP)
		for i := 0; i < c.beginP; i++ {
			fmt.Fscan(reader, &(avialable[i]))
		}
		submits := make([]Submit, 0)
		for {
			s := getLine(r)
			if s == "ENDS" {
				break
			}
			var t time.Time
			var task, p int
			fmt.Sscan(s, &t)
			fmt.Sscan(getLine(r), &task)
			fmt.Sscan(getLine(r), &p)
			submits = append(submits, Submit{t, task, p})
		}
		c.users[name] = User{
			name,
			password,
			points,
			avialable,
			submits,
		}
	}
	line := getLine(r)
	scores := strings.Split(line, " ")
	C.scoring = make([]int, len(scores))
	for i := 0; i < len(scores); i++ {
		fmt.Sscan(scores[i], &(C.scoring[i]))
	}
}

func (c *Contest) write(w *bufio.Writer) {
	fmt.Fprintln(w, c.start)
	fmt.Fprintln(w, c.beginP)
	for _, p := range c.problemset {
		(&p).write(w)
	}
	fmt.Fprintln(w, "ENDP")
	for _, u := range c.users {
		fmt.Fprintln(w, u.name)
		fmt.Fprintln(w, u.password)
		fmt.Fprintln(w, u.points)
		for i, a := range u.avialable {
			if i > 0 {
				fmt.Fprintf(w, " ")
			}
			fmt.Fprint(w, a)
		}
		fmt.Fprintln(w, "")
		for _, s := range u.submits {
			fmt.Fprintf(w, "%v\n%v\n%v\n", s.t, s.task, s.points)
		}
		fmt.Fprintln(w, "ENDS")
	}
	fmt.Fprintln(w, "ENDU")
	for i, sc := range c.scoring {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, sc)
	}
	fmt.Fprintln(w)
}

func (s Scoreboard) write(w *bufio.Writer) {
	fmt.Fprintln(w, s.String())
}

func getLine(r *bufio.Reader) string {
	line, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return line[:len(line)-1]
}
