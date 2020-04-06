package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//User is
type User struct {
	Username string
	Password string
}

//DataUser is database
type DataUser struct {
	ID       int
	Username string
	Password string
}

//Post is posts
type Post struct {
	ID       int
	Title    string
	Subtitle string
}

//Cpost sda
type Cpost struct {
	Subject string
	Message string
}

//veri tabanı error kontol fonk.
func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

func run(response http.ResponseWriter, requst *http.Request) {
	response.Write([]byte("runing"))

}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("sqlite3", "db/blog.db")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select * from posts")
	checkError(err)
	var tempPost Post

	for rows.Next() {

		err = rows.Scan(&tempPost.ID, &tempPost.Title, &tempPost.Subtitle)
		if err != nil {
			// handle this error
			panic(err)
		}

	}
	json.NewEncoder(w).Encode(tempPost)
}
func login(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "db/blog.db")
	if err != nil {
		panic(err)
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	b := string(reqBody)
	var user User
	json.Unmarshal([]byte(b), &user)
	name := user.Username
	pass := user.Password
	fmt.Print(name)
	//veri tabanı verileri çekme
	rows, err := db.Query("select * from users")
	checkError(err)
	for rows.Next() {
		var tempUser DataUser
		err = rows.Scan(&tempUser.ID, &tempUser.Username, &tempUser.Password)
		checkError(err)
		if tempUser.Username == name && tempUser.Password == pass {
			w.Write([]byte("true"))
		} else {

			w.Write([]byte("false"))
		}

	}

	//fmt.Println(getUser(db, name, pass))

}

func addpost(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "db/blog.db")
	if err != nil {
		panic(err)
	}
	reqBody2, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	a := string(reqBody2)
	var Cpost Cpost
	json.Unmarshal([]byte(a), &Cpost)
	subject := Cpost.Subject
	message := Cpost.Message

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO posts (title, subject) VALUES (?, ?)")
	stmt.Exec(subject, message)

	tx.Commit()

}

/*func getUser(db *sql.DB, name1 string, paswd string) bool {
	rows, err := db.Query("select * from users")
	checkError(err)
	for rows.Next() {
		var tempUser DataUser
		err = rows.Scan(&tempUser.ID, &tempUser.Username, &tempUser.Password)
		checkError(err)
		if tempUser.Username == name1 && tempUser.Password == paswd {
			return true
		}
	}
	return false
}*/
func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/run", run).Methods("GET")
	router.HandleFunc("/posts", posts).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/addpost", addpost).Methods("POST")

	http.ListenAndServe(":8070", handlers.CORS(headers, methods, origins)(router))

}
