package main

import (
	"crypto/sha1"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var t *template.Template

type FileSystem struct {
	fs http.FileSystem
}

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/index", index)
	http.Handle("/favico", http.NotFoundHandler())
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	if r.Method == http.MethodPost {
		multiFile, fileHeader, err := r.FormFile("filename")
		if err != nil {
			log.Println(err)
		}
		defer multiFile.Close()


		log.Println("The filename is: ", fileHeader.Filename)

		h := sha1.New()
		log.Println("The value of h is: ", h)
		io.Copy(h, multiFile)

		workingDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		workingDirPath := filepath.Join(workingDir, "public", "pics", fileHeader.Filename)
		dir, err := os.Create(workingDirPath)
		if err != nil {
			log.Println("ERROR: ", err)
		}
		defer dir.Close()
		log.Println("The working path is: ", workingDirPath, " has been created.")
		multiFile.Seek(0, 0)
		io.Copy(dir, multiFile)

		cookie = appendValue(w, cookie, fileHeader.Filename)
		xd := strings.Split(cookie.Value, "|")
		t.ExecuteTemplate(w, "retrieveemployee.gohtml", xd[1:])

	} else {
		t.ExecuteTemplate(w, "retrieveemployee.gohtml", nil)
	}
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:       "session",
			Value:      id.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fileName string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fileName) {
		s += "|" + fileName
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}