package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>[<a href=\"/show/\">Ukáž príklad</a>]</p>"+
		"<p>[<a href=\"/show/\">Submit</a>]</p>"+
		"<p>[<a href=\"/scoreboard/\">Ukáž výsledkovku</a>]</p>")
}

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
	if len(values) < 3 {
		fmt.Fprint(w, "<p>[<a href=\"/show/\">Back</a>]</p>")
		http.Error(w, NotEnoughDataError{}.Error(), http.StatusInternalServerError)
		return
	}
	name, password := values[0], values[1]
	_, err := fmt.Sscan(values[2], &num)
	fmt.Println("User", name, "password", password, "to show problem", num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statement, image, err := showp(name, password, num)

	if err != nil {
		fmt.Fprint(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		User, Password, Statement string
		Image                     string
		Id                        int
	}{
		name,
		password,
		statement,
		image,
		num,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if !active {
		fmt.Fprintln(w, "No contest is running!")
		return
	}
	var num int
	url := r.URL.Path[len("/submit/"):]
	fmt.Println("Submit| url: ", url, "cela: ", r.URL.Path)
	values := strings.Split(url, "/")
	if len(values) < 3 {
		fmt.Fprint(w, "<p>[<a href=\"/\">Back</a>]</p>")
		http.Error(w, NotEnoughDataError{}.Error(), http.StatusInternalServerError)
		return
	}
	name, password := values[0], values[1]
	_, err := fmt.Sscan(values[2], &num)
	fmt.Println("User", name, "password", password, "to show problem", num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("submit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statement, image, err := showp(name, password, num)

	if err != nil {
		fmt.Fprint(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		User, Password, Statement string
		Image                     string
		Id                        int
	}{
		name,
		password,
		statement,
		image,
		num,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func evaluateHandler(w http.ResponseWriter, r *http.Request) {
	if !active {
		fmt.Fprintln(w, "No contest is running!")
		return
	}
	fmt.Fprint(w, "<p>[<a href=\"/\">Back</a>]</p>")
	var num int
	url := r.URL.Path[len("/evaluate/"):]
	values := strings.Split(url, "/")
	if len(values) < 3 {
		fmt.Fprint(w, "<p>[<a href=\"/\">Back</a>]</p>")
		http.Error(w, NotEnoughDataError{}.Error(), http.StatusInternalServerError)
		return
	}
	name, password := values[0], values[1]
	_, err := fmt.Sscan(values[2], &num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	answer := r.FormValue("answer")

	points, err := C.submit(name, password, num, answer)
	if err != nil {
		fmt.Println("Tu je problem?")
		fmt.Fprint(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Dostal si %v bodov.", points)
}

func scoreboardHandler(w http.ResponseWriter, r *http.Request) {
	if !active {
		fmt.Fprintln(w, "No contest is running!")
		return
	}

	t, err := template.ParseFiles("scoreboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, struct {
		Scores Scoreboard
	}{
		C.getScoreboard(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/"):]
	http.ServeFile(w, r, path)
}
