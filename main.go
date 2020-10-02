package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var stat int

type MyMux struct {
}

type UsersInfo struct {
	Users UserInfo
}
type UserInfo struct {
	ID        int
	Username  string
	Password  string
	Privilege string
}

func createdata(w http.ResponseWriter, r *http.Request) {
	stat = 0
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/input.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])

		if len(r.Form.Get("username")) < 5 {

			fmt.Println("Error, username tidak boleh kurang dari 5 karakter")
		} else if len(r.Form.Get("password")) < 8 {
			fmt.Println("Error, password tidak boleh kurang dari 8 karakter")
		} else {
			stat = 1
		}

		if stat == 1 {
			db, dberr := sql.Open("sqlite3", "./login1.db")
			checkErr(dberr)

			stmt, dberr := db.Prepare("insert into `userinfo` (`uid`,`username` , `password`) values (null, ?, ?)")
			checkErr(dberr)

			res, dberr := stmt.Exec(r.Form.Get("username"), r.Form.Get("password"))
			checkErr(dberr)

			id, dberr := res.LastInsertId()
			checkErr(dberr)

			fmt.Println(id)

			http.Redirect(w, r, "http://localhost:9096/show", 303)
		} else {
			http.Redirect(w, r, "http://localhost:9096/input", 303)
		}

	}

}

func readdata(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.ParseFiles("./templates/showdata.gtpl")
	check(err)
	db, dberr := sql.Open("sqlite3", "./login1.db")
	checkErr(dberr)
	rows, dberr := db.Query("SELECT * FROM userinfo where uid=?", 2)
	checkErr(dberr)
	var uid, privilege int
	var username, password string
	var users UsersInfo

	for rows.Next() {
		dberr = rows.Scan(&uid, &username, &password, &privilege)
		checkErr(dberr)

		users.Users.ID = uid
		users.Users.Username = username
		users.Users.Password = password
	}

	rows.Close()

	err = t.Execute(w, users)
	check(err)
}

func editdata(w http.ResponseWriter, r *http.Request) {
	stat = 0
	fmt.Println("Method: ", r.Method)
	if r.Method == "POST" {
		t, _ := template.ParseFiles("./templates/edit.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])

		if len(r.Form.Get("username")) < 5 {

			fmt.Println("Error, username tidak boleh kurang dari 5 karakter")
		} else if len(r.Form.Get("password")) < 8 {
			fmt.Println("Error, password tidak boleh kurang dari 8 karakter")
		} else {
			stat = 1
		}

		if stat == 1 {
			db, dberr := sql.Open("sqlite3", "./login1.db")
			checkErr(dberr)

			stmt, dberr := db.Prepare("update `userinfo` set(`username` , `password`) values (?, ?) where uid=?")
			checkErr(dberr)

			res, dberr := stmt.Exec(r.Form.Get("username"), r.Form.Get("password"), 2)
			checkErr(dberr)

			id, dberr := res.LastInsertId()
			checkErr(dberr)

			fmt.Println(id)

			http.Redirect(w, r, "http://localhost:9096/show", 303)
		} else {
			http.Redirect(w, r, "http://localhost:9096/input", 303)
		}

	}

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/input" {
		createdata(w, r)
		return
	} else if r.URL.Path == "/show" {
		readdata(w, r)
		return
	} else if r.URL.Path == "/edit" {
		editdata(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":9096", mux)
	if err != nil {
		log.Fatal("Error running service:", err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
