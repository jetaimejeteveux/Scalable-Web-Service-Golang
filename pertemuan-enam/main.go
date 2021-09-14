package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var studentData = []student{
	{
		ID:    "1",
		Name:  "Ali",
		Grade: 6,
	},
	{
		ID:    "2",
		Name:  "Lia",
		Grade: 7,
	},
	{
		ID:    "3",
		Name:  "Ila",
		Grade: 8,
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	http.HandleFunc("/index", index)
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("server runninn port 8080")

	http.HandleFunc("/golang", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := map[string]string{
			"Firstname": "ila",
			"lastname":  "lia",
		}
		t.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world dari index")
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		res, err := json.Marshal(studentData)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		w.Write(res)
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		id := r.FormValue("id")
		for _, v := range studentData {
			if v.ID == id {
				res, err := json.Marshal(v)
				if err != nil {
					http.Error(w, "Bad Request", http.StatusBadRequest)
					return
				}
				w.Write(res)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
