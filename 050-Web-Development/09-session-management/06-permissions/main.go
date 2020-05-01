package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var t *template.Template

type User struct {
	Firstname string
	Lastname string
	Username string
	Password []byte
	Role string
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
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/health", health)
	http.HandleFunc("/instance", instance)

	http.ListenAndServe(":80", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if len(userDB) == 0 {
		log.Println("No entry in DB, redirecting to signup")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		log.Println("Checking for cookie")
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Println("No Session active at the moment")
			username := r.FormValue("username")
			password := r.FormValue("password")
			for key,_ := range userDB {
				fmt.Println(username, password, userDB[key].Username, bcrypt.CompareHashAndPassword(userDB[key].Password, []byte(password)), sessionDB)
				log.Println("Validating the existance of the user in the DB")
				log.Println("user matches:", userDB[key].Username == username , "password matches: ", bcrypt.CompareHashAndPassword(userDB[key].Password, []byte(password)) == nil)
				if userDB[key].Username == username && bcrypt.CompareHashAndPassword(userDB[key].Password, []byte(password)) == nil {
					log.Println("username: ", username, " exists in the DB and the password matches")
					id, _ := uuid.NewV4()
					cookie := &http.Cookie{
						Name:       "session",
						Value:      id.String(),
					}
					http.SetCookie(w, cookie)
					sessionDB[cookie.Value] = username
					Userdetails := userDB[username]
					t.ExecuteTemplate(w, "login.gohtml", Userdetails)
					return
				} else if userDB[key].Username == username && bcrypt.CompareHashAndPassword(userDB[key].Password, []byte(password)) != nil {
					log.Println("Password incorrect, redirecting to login page again")
					t.ExecuteTemplate(w, "login.gohtml", nil)
					return
				} else {
					continue
				}
			}
			log.Println("Password incorrect, redirecting to login page again")
			t.ExecuteTemplate(w, "login.gohtml", nil)
			return
		} else {
			log.Println("Session active at the moment")
			username := sessionDB[cookie.Value]
			Userdetails := userDB[username]
			t.ExecuteTemplate(w, "login.gohtml", Userdetails)
		}
	}
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

func register(firstname, lastname, username, password, role string) bool {
	var user User

	user.Username = username
	user.Firstname = firstname
	user.Lastname = lastname
	user.Role = role
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
		role := r.FormValue("role")

		if register(firstname, lastname, username, password,role) {
			fmt.Println("Registration Successful")
			t.ExecuteTemplate(w, "login.gohtml", nil)
			return
		}
		t.ExecuteTemplate(w, "signup.gohtml", nil)
		return
	}

	t.ExecuteTemplate(w, "signup.gohtml", nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "admin.gohtml", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("status", "200")
	io.WriteString(w, "GOOD")
}

func instance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("status", "200")
	io.WriteString(w, "This is from VM2")
}
