package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var t *template.Template

type User struct {
	Firstname string
	Lastname string
	Userid string
	Password string
}

var sessionDB = make(map[string]string)
var userDB = make(map[string]User)

func init() {
	t = template.Must(t.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)

	http.ListenAndServe(":8080", nil)
}


func login(w http.ResponseWriter, r *http.Request) {
	if len(userDB) != 0 {
		fmt.Println("userDB has elements")
		if r.Method == http.MethodPost {
			cookie, err := r.Cookie("session")
			if err != nil {
				fmt.Println("Cookie doesn't exist, hence logging in..")
				fmt.Println("Receiving Username and Password")

				username := r.FormValue("userid")
				password := r.FormValue("password")

				for key,_ := range userDB {
					switch {
					case userDB[key].Userid == username && userDB[key].Password == password:
						id, _ := uuid.NewV4()
						cookie := &http.Cookie{
							Name:       "session",
							Value:      id.String(),
						}

						sessionDB[id.String()] = username

						http.SetCookie(w, cookie)
						Userdetails := userDB[username]
						t.ExecuteTemplate(w, "login.gohtml", Userdetails)

					case userDB[key].Userid == username && userDB[key].Password != password:
						fmt.Println("Wrong Password")
						t.ExecuteTemplate(w, "login.gohtml", nil)

					case userDB[key].Userid != username:
						fmt.Println("userid: ", username, " doesn't exist")
						t.ExecuteTemplate(w, "login.gohtml", nil)

					case sessionDB[cookie.Value] == username:
						Userdetails := userDB[username]
						t.ExecuteTemplate(w, "login.gohtml", Userdetails)

					case sessionDB[cookie.Value] != username:
						for key,_ := range sessionDB {
							if key == cookie.Value {
								http.SetCookie(w, cookie)
								Userdetails := userDB[username]
								t.ExecuteTemplate(w, "login.gohtml", Userdetails)
							}
						}
					}
				}
			} else {
				fmt.Println("Session already exists")
				username := sessionDB[cookie.Value]
				Userdetails := userDB[username]
				t.ExecuteTemplate(w, "login.gohtml", Userdetails)
			}

		} else {
			t.ExecuteTemplate(w, "login.gohtml", nil)
		}
	} else {
		fmt.Println("userDB is empty")
		t.ExecuteTemplate(w, "signup.gohtml", nil)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatalln("Cookie has not been set")
	}

	delete(sessionDB, cookie.Value)
	cookie.MaxAge = -1
	t.ExecuteTemplate(w, "logout.gohtml", nil)
}

func register(firstname, lastname, userid, password string) {
	fmt.Printf("Registering user into the DB, with username: %v, password: %v\n", userid, password)
	var user User

	user.Firstname = firstname
	user.Password = password
	user.Userid = userid
	user.Lastname = lastname

	userDB[user.Userid] = user
	fmt.Println(userDB)
	fmt.Println("Entry for the userid: ", userid, " inserted into the database. You can login now")

}

func signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("Capturing the Form values")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		userid := r.FormValue("username")
		password := r.FormValue("password")

		if len(userDB) == 0 {
			fmt.Println("User DB is empty")
			register(firstname, lastname, userid, password)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
		for key, _ := range userDB {
				fmt.Println("Checking DB")
				if userDB[key].Userid == userid {
					fmt.Println("user id: ", userid, " already exists, please try again..")
					http.Redirect(w, r, "/signup", http.StatusSeeOther)
					return
				} else {
					fmt.Println("User not found, registering")
					register(firstname, lastname, userid, password)
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			}
		}
	} else {
		t.ExecuteTemplate(w, "signup.gohtml", nil)
	}
}
