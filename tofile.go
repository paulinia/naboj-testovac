package main

import (
	"bufio"
	"fmt"
)

func (d *Database) writeToFile(w *bufio.Writer) {
	for _, p := range d.problems {
		fmt.Fprintln(w, p.statement)
		fmt.Fprintln(w, "ENDS")
		fmt.Fprintln(w, p.result)
		fmt.Fprintln(w, p.level)
		fmt.Fprintln(w, p.typ)
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
		d.addProblem(Problem{
			statement,
			res,
			lev,
			getLine(r),
		})
	}
}

func (c *Contest) readContest(r *bufio.Reader) {

}

func (c *Contest) writeContest(w *bufio.Writer) {
	fmt.Fprintln(w, c.n)
	fmt.Fprintln(w, c.beginP)
}

func getLine(r *bufio.Reader) string {
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return line[:len(line)-1]
}
