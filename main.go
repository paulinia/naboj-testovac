package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	D = Database{
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

	fmt.Println("Type 'NEW' to enter new problem; " +
		"'GEN' to generate new contest; " +
		"'START' to start/resume a contest; " +
		"'END' to end a contest; " +
		"'USER' to add a user; " +
		"'QUIT' to quit.")

	r = bufio.NewReader(os.Stdin)
	active, generated = false, false
	defer func() {
		if generated {
			fw, _ := os.Create("contest.txt")
			w := bufio.NewWriter(fw)
			C.write(w)
			w.Flush()
			fw.Close()
		}
	}()

	fc, er := os.Open("contest.txt")
	if er == nil {
		C.read(bufio.NewReader(fc))
		generated = true
		defer fc.Close()
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/scoreboard/", scoreboardHandler)
	http.HandleFunc("/show/", showHandler)
	http.HandleFunc("/submit/", submitHandler)
	http.HandleFunc("/evaluate/", evaluateHandler)
	http.HandleFunc("/problem/", problemHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/image/", imageHandler)
	go http.ListenAndServe(":8000", nil)

loop:
	for {
		switch getLine(r) {
		case "NEW":
			newProblem()
		case "GEN":
			if active {
				fmt.Printf("A contest is already running!\n")
				break
			}
			if generated {
				sc := C
				defer func() {
					fw, _ := os.Create(fmt.Sprintf("%v.sb", sc.start))
					w := bufio.NewWriter(fw)
					sc.getScoreboard().write(w)
					w.Flush()
					fw.Close()
				}()
			}
			generate()
		case "QUIT":
			break loop
		case "START":
			if !generated {
				fmt.Printf("There's no contest to start.\n")
				break
			}
			if active {
				fmt.Printf("A contest is already running!")
				break
			}
			C.begin()
			active = true
		case "END":
			if !active {
				fmt.Printf("There's no contest running.\n")
				break
			}
			C.end()
			active = false
			fmt.Printf(C.getScoreboard().String())
			sc := C
			defer func() {
				fw, _ := os.Create(fmt.Sprintf("%v.sb", sc.start))
				w := bufio.NewWriter(fw)
				sc.getScoreboard().write(w)
				w.Flush()
				fw.Close()
			}()
		case "USER":
			if !generated {
				fmt.Printf("There's no contest running\n")
				break
			}
			addUser()
		default:
			fmt.Println("Type 'NEW' to enter new problem; " +
				"'GEN' to generate new contest; " +
				"'START' to start/resume a contest; " +
				"'END' to end a contest; " +
				"'USER' to add a user; " +
				"'QUIT' to quit.")
		}
	}

}
