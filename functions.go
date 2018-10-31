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
	typ := getLine(r)
	fmt.Println("Zadaj adresu obrazku (alebo prazdny string ak ziadny): ")
	D.addProblem(Problem{
		statement,
		res,
		num,
		typ,
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
			if len(scoring) < 1 {
				scoring = append(scoring, 1)
			}
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
	if err := C.addUser(name, password); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Succesful")
}

func submit(name, password, sol string, task int) (int, error) {
	fmt.Printf("Riešenie (viacero čísel oddeľujte 'a') : ")
	p, err := C.submit(name, password, task, sol)
	return p, err
}

func showp(name, password string, task int) (string, string, error) {
	return C.show(name, password, task)
}
