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
		w := bufio.NewWriter(fw)
		D.writeToFile(w)
		w.Flush()
		fw.Close()
	}()

	r := bufio.NewReader(os.Stdin)
	fmt.Println("Type 'NEW' to enter new problem; " +
		"'GEN' to generate new contest; " +
		"'START' to start/resume a contest; " +
		"'END' to end a contest; " +
		"'SHOWS' to show a scoreboard; " +
		"'USER' to add a user; " +
		"'SHOWP' to show a problem; " +
		"'SUBMIT' to submit a solution and " +
		"'QUIT' to quit.")

	var C Contest
	active, generated := false, false
	defer func() {
		if C.n > 0 {
			fw, _ := os.Create("constest.txt")
			w := bufio.NewWriter(fw)
			C.write(w)
			w.Flush()
			fw.Close()
		}
	}()

loop:
	for {
		switch getLine(r) {
		case "NEW":
			fmt.Printf("Zadaj zadanie prikladu. Ukonci ho s riadkom 'ENDS'.: ")
			statement := ""
			line := getLine(r)
			for {
				if line == "ENDS" {
					break
				}
				statement += line
				line = getLine(r)
			}
			fmt.Printf("Zadaj riešenie príkladu: ")
			var num int
			res := toResult(getLine(r))
			fmt.Printf("Zadaj číslo príkladu: ")
			fmt.Sscan(getLine(r), &num)
			fmt.Printf("Zadaj typ príkladu [MAT/FYZ]: ")
			D.addProblem(Problem{
				statement,
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
			active = false

			sc := C
			defer func() {
				fw, _ := os.Create(fmt.Sprintf("%v.sb", sc.start))
				w := bufio.NewWriter(fw)
				sc.getScoreboard().write(w)
				w.Flush()
				fw.Close()
			}()
			fmt.Printf(C.getScoreboard().String())
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
		case "SHOWP":
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
		case "SHOWS":
			if !active {
				fmt.Printf("No contest running!\n")
				break
			}
			fmt.Printf(C.getScoreboard().String())
		default:
			fmt.Println("Type 'NEW' to enter new problem; " +
				"'GEN' to generate new contest; " +
				"'START' to start/resume a contest; " +
				"'END' to end a contest; " +
				"'SHOWS' to show a scoreboard; " +
				"'USER' to add a user; " +
				"'SHOWP' to show a problem; " +
				"'SUBMIT' to submit a solution and " +
				"'QUIT' to quit.")
		}
	}

}
