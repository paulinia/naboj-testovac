package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func showHandler(w http.ResponseWriter, r *http.Request) {
	if !active {
		fmt.Fprintln(w, "No contest is running!")
		return
	}
	fmt.Fprintf(w, "<h1>Zobraz príklad</h1>\n"+
		"\n"+
		"<form action=\"/problem/\" method=\"POST\">\n"+
		"Login: \n"+
		"<div><input type=\"text\" name=\"name\" value = \"\"></div>\n"+
		"Heslo:\n"+
		"<div><input type=\"text\" name=\"password\" value = \"*****\"></div>\n"+
		"Číslo príkladu:\n"+
		"<div><input type=\"text\" name=\"number\"></div>\n"+
		"<div><input type=\"submit\" value=\"Ukáž\"></div>\n"+
		"</form>\n"+
		"")
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	if !active {
		fmt.Fprintln(w, "No contest is running!")
		return
	}
	fmt.Printf("WTF toto preco nefunguje?")
	name := r.FormValue("name")
	password := r.FormValue("password")
	num := r.FormValue("number")
	fmt.Println("/view/" + name + "/" + password + "/" + num)
	http.Redirect(w, r, "/view/"+name+"/"+password+"/"+num, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if !active {
		fmt.Fprintln(w, "No contest is running!")
		return
	}
	var num int
	url := r.URL.Path[len("/save/"):]
	fmt.Println("url: ", url)
	values := strings.Split(url, "/")
	name, password := values[0], values[1]
	_, err := fmt.Sscan(values[2], &num)
	fmt.Println("User", name, "password", password, "to show problem", num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Necrashlo to tuna")
	t, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Sparsovane")

	statement, err := showp(name, password, num)

	if err != nil {
		fmt.Fprint(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		User, Password, Statement string
		Id                        int
	}{
		name,
		password,
		statement,
		num,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
