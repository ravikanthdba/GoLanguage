package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var t *template.Template

type User struct {
	Firstname string
	Lastname string
	Username string
	Password []byte
}

var sessionDB = make(map[string]string)
var userDB = make(map[string]User)

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)

	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if len(userDB) == 0 {
		fmt.Println("User Database has no elements, hence going to the signup page")
		t.ExecuteTemplate(w, "signup.gohtml", nil)
		return
	} else {
		if r.Method == http.MethodPost {
			cookie, err := r.Cookie("session")
			fmt.Println(cookie)
			username := r.FormValue("username")
			password := r.FormValue("password")


			if err != nil {
				// Cookie doesn't exist, Validate the user, generate the cookie, add the session details to sessionDB Database
				for key,_ := range userDB {
					switch {
					case userDB[key].Username == username &&  bcrypt.CompareHashAndPassword(userDB[key].Password, []byte(password)) == nil:
						fmt.Println("Credentials are correct, logging in")
						id, _ := uuid.NewV4()
						cookie := &http.Cookie{
							Name:       "session",
							Value:      id.String(),
						}
						http.SetCookie(w, cookie)
						sessionDB[cookie.Value] = username
						fmt.Println("Username is: ", username)
						Userdetails := userDB[username]
						fmt.Println(Userdetails)
						t.ExecuteTemplate(w, "login.gohtml", Userdetails)
						return
					case userDB[key].Username == username && bcrypt.CompareHashAndPassword(userDB[key].Password, []byte(password)) != nil:
						fmt.Println("Password incorrect, please try logging in with the right password")
						t.ExecuteTemplate(w, "login.gohtml", nil)
					case userDB[key].Username != username:
						fmt.Println("Username incorrect, please try logging in with the right username")
						t.ExecuteTemplate(w, "login.gohtml", nil)
					}
				}

				return
			} else {
				Userdetails := userDB[sessionDB[cookie.Value]]
				fmt.Println(Userdetails)
				t.ExecuteTemplate(w, "login.gohtml", Userdetails)
				return
			}
		}
	}
	t.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	cookie := &http.Cookie{
		Name:       "session",
		Value:      "",
		MaxAge:     -1,
	}
	http.SetCookie(w, cookie)
	delete(sessionDB, c.Value)
	t.ExecuteTemplate(w, "logout.gohtml", nil)
}

func register(firstname, lastname, username, password string) bool {
	var user User

	user.Username = username
	user.Firstname = firstname
	user.Lastname = lastname
	encryptedPassword,_ := bcrypt.GenerateFromPassword([]byte(password), 4)
	user.Password = encryptedPassword

	fmt.Println(user)

	if len(userDB) == 0 {
		fmt.Println("User DB is blank, hence registering immediately")
		userDB[username] = user
		return true
	} else {
		fmt.Println("Checking if user name already exists")
		for key,_ := range userDB {
			if userDB[key].Username == username {
				fmt.Println("Username is already used, please retry")
				return false
			} else {
				fmt.Println("Username not exists in the DB, registering")
				userDB[username] = user
				return true
			}
		}
	}
	return false
}

func signup(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err == nil {
			fmt.Println("Found Cookie, Setting cookie value to -1")
			c := &http.Cookie{
				Name:       "session",
				Value:      "",
				MaxAge:     -1,
			}
			http.SetCookie(w, c)
			delete(sessionDB, cookie.Value)
		}

		if r.Method == http.MethodPost {
			firstname := r.FormValue("firstname")
			lastname := r.FormValue("lastname")
			username := r.FormValue("username")
			password := r.FormValue("password")

			if register(firstname, lastname, username, password) {
				fmt.Println("Registration Successful")
				t.ExecuteTemplate(w, "login.gohtml", nil)
				return
			}
			t.ExecuteTemplate(w, "signup.gohtml", nil)
			return
		}

		t.ExecuteTemplate(w, "signup.gohtml", nil)
}
