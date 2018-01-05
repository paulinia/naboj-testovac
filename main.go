package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println("D je ", D)
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
			generated = true

			C = generateContest(n, bn, []string{typ}, &D)
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
			active = false
			generated = false
			fmt.Println(C.getScoreboard())
		case "USER":
			if !active {
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
			}
			fmt.Printf("Username: ")
			name := getLine(r)
			fmt.Printf("Password ")
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
		default:
			fmt.Println("Type 'NEW' to enter new problem; " +
				"'GEN' to generate new contest; " +
				"'START' to start/resume a contest; " +
				"'END' to end a contest; " +
				"'USER' to add a user " +
				"'SUBMIT' to submit a solution and " +
				"'QUIT' to quit.")
		}
	}

}
