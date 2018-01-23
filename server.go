package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func showHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Zobraz príklad</h1>\n"+
		"\n"+
		"<form action=\"/problem/\" method=\"POST\">\n"+
		"Login: \n"+
		"<div><input type=\"text\" name=\"author\"></div>\n"+
		"Heslo:\n"+
		"<div><input type=\"text\" name=\"password\"></div>\n"+
		"Číslo príkladu:\n"+
		"<div><input type=\"text\" name=\"number\"></div>\n"+
		"<div><input type=\"submit\" value=\"Ukáž\"></div>\n"+
		"</form>\n"+
		"")
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	num := r.FormValue("number")
	http.Redirect(w, r, "/view/"+name+"/"+password+"/"+num, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var name, password string
	var num int
	url := r.URL.Path[len("/save/"):]
	_, err := fmt.Sscanf(url, "%v/%v/%v", &name, &password, &num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statement, er := showp(name, password, num)

	if er != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	}
}
