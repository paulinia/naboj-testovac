package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	D := Database{
		0,
		0,
		0,
		make([]Problem, 0),
	}
	f, er := os.Open("database.txt")
	if er == nil {
		fmt.Println("er nie je")
		D.readDatabase(bufio.NewReader(f))
		defer f.Close()
	}

	defer func() {
		fw, _ := os.Create("database.txt")
		D.writeToFile(bufio.NewWriter(fw))
		fw.Close()
	}()

	r := bufio.NewReader(os.Stdin)
	fmt.Println("Type 'NEW' to enter new problem; " +
		"'GEN' to generate new contest; " +
		"'START' to start/resume a contest; " +
		"'END' to end a contest; " +
		"'USER' to add a user " +
		"'SHOW' to show a problem; " +
		"'SUBMIT' to submit a solution and " +
		"'QUIT' to quit.")

	var C Contest
	active, generated := false, false

loop:
	for {
		switch getLine(r) {
		case "NEW":
			line := getLine(r)
			var num int
			res := toResult(getLine(r))
			fmt.Sscan(getLine(r), &num)
			D.addProblem(Problem{
				line,
				res,
				num,
				getLine(r),
			})
		case "GEN":
			if active {
				fmt.Printf("A contest is already running!\n")
				break
			}
			fmt.Printf("Napíš s koľkými príkladmi je náboj a " +
				"s koľkými začínajú a 'MAT' / 'FYZ' podľa typu náboja : ")
			var n, bn int
			var typ string
			_, err := fmt.Sscan(getLine(r), &n, &bn, &typ)
			if err != nil {
				break
			}
			fmt.Printf("Napíš bodovanie (počet bodov na prvy, druhý tretí atď submit a za ostatné): ")
			scoring := make([]int, 0)
			reader := strings.NewReader(getLine(r))

			for {
				var num int
				_, err := fmt.Fscan(reader, &num)
				if err == io.EOF {
					generated = true
					C = generateContest(n, bn, []string{typ}, scoring, &D)
					break
				}
				if err != nil {
					fmt.Printf("Wrong format. Try again.\n")
					break
				}
				scoring = append(scoring, num)
			}
		case "QUIT":
			break loop
		case "START":
			if !generated {
				fmt.Printf("There's no contest to start.\n")
				break
			}
			C.begin()
			active = true
		case "END":
			if !active {
				fmt.Printf("There's no contest to end.\n")
				break
			}
			C.end()
			fmt.Printf("Co toto preboha robi?\n")
			active = false
			generated = false
			C.getScoreboard().show()
		case "USER":
			if !generated {
				fmt.Printf("There's no contest running\n")
				break
			}
			fmt.Printf("Username: ")
			name := getLine(r)
			fmt.Printf("Password: ")
			password := getLine(r)
			C.addUser(name, password)
		case "SUBMIT":
			if !active {
				fmt.Printf("There's no contest running\n")
				break
			}
			fmt.Printf("Username: ")
			name := getLine(r)
			fmt.Printf("Password: ")
			password := getLine(r)
			fmt.Printf("Task id: ")
			var task int
			fmt.Sscan(getLine(r), &task)
			fmt.Printf("Riešenie (viacero čísel oddeľujte 'a') : ")
			sol := getLine(r)
			p, err := C.submit(name, password, task, sol)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Dostal si", p, "bodov.")
			}
		case "SHOW":
			if !active {
				fmt.Printf("There's no contest running\n")
				break
			}
			fmt.Printf("Username: ")
			name := getLine(r)
			fmt.Printf("Password: ")
			password := getLine(r)
			fmt.Printf("Task id: ")
			var task int
			fmt.Sscan(getLine(r), &task)
			s, er := C.show(name, password, task)
			if er != nil {
				fmt.Println(er.Error())
			} else {
				fmt.Println(s)
			}
		default:
			fmt.Println("Type 'NEW' to enter new problem; " +
				"'GEN' to generate new contest; " +
				"'START' to start/resume a contest; " +
				"'END' to end a contest; " +
				"'USER' to add a user; " +
				"'SHOW' to show a problem; " +
				"'SUBMIT' to submit a solution and " +
				"'QUIT' to quit.")
		}
	}

}
