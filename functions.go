package main

import (
	"fmt"
	"io"
	"strings"
)

func newProblem() {
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
}

func generate() {
	fmt.Printf("Napíš s koľkými príkladmi je náboj a " +
		"s koľkými začínajú a 'MAT' / 'FYZ' podľa typu náboja : ")
	var n, bn int
	var typ string
	_, err := fmt.Sscan(getLine(r), &n, &bn, &typ)
	if err != nil {
		return
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
}

func addUser() {
	fmt.Printf("Username: ")
	name := getLine(r)
	fmt.Printf("Password: ")
	password := getLine(r)
	C.addUser(name, password)
}

func submit() {
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
}

func showp(name, password string, task int) (string, error) {
	fmt.Sscan(getLine(r), &task)
	s, er := C.show(name, password, task)
	return s, er
}
