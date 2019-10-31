package main

import (
    "net/http"

    "database/sql"
    "encoding/json"
    "net/url"
    "encoding/xml"
    "io/ioutil"

    _ "github.com/mattn/go-sqlite3"
    "github.com/urfave/negroni"
    "github.com/yosssi/ace"
    gmux "github.com/gorilla/mux"
)

type Book struct {
	PK int
	Title string
	Author string
	Classification string
}

type Page struct {
    Books []Book
}


var db *sql.DB;

func verifyDatabase(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    if err := db.Ping(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    next(w, r)
}

func ClassifyAPI(url string) ([]byte, error) {
    var resp *http.Response
    var err error


    resp, err = http.Get(url)
    if err != nil {
        return []byte{}, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}

type SearchResult struct {
    Title string `xml:"title,attr"`
    Author string `xml:"author,attr"`
    Year string `xml:"hyr,attr"`
    ID string `xml:"owi,attr"`
}

type ClassifySearchResponse struct {
    Results []SearchResult `xml:"works>work"`
}

func search(query string) ([]SearchResult, error) {
    var c ClassifySearchResponse

    body, err := ClassifyAPI("http://classify.oclc.org/classify2/Classify?summary=true&title=" + url.QueryEscape(query))
    if err != nil {
        return []SearchResult{}, err
    }

    err = xml.Unmarshal(body, &c)
    return c.Results, err
}

type ClassifyBookResponse struct {
    BookData struct {
        Title string `xml:"title,attr"`
        Author string `xml:"author,attr"`
        ID string `xml:"owi,attr"`
    } `xml:"work"`
    Classification struct {
        MostPopular string `xml:"sfa,attr"`
    } `xml:"recommendations>ddc>mostPopular"`
}


func find(id string) (ClassifyBookResponse, error) {
    var c ClassifyBookResponse

    body, err := ClassifyAPI("http://classify.oclc.org/classify2/Classify?summary=true&owi=" + url.QueryEscape(id))
    if err != nil {
        return ClassifyBookResponse{}, err
    }

    err = xml.Unmarshal(body, &c)
    return c, err
}

func main() {

    mux := gmux.NewRouter()

    db, _ = sql.Open("sqlite3", "dev.db")



    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        template, err := ace.Load("templates/index", "", nil)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        p := Page {Books: []Book{}}
        rows, _ := db.Query("select pk, title, author, classification from books")
        for rows.Next() {
        	var b Book
        	rows.Scan(&b.PK, &b.Title, &b.Author, b.Classification)
        	p.Books = append(p.Books, b)
        }

        _ = template.Execute(w, p)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

    }).Methods("GET")





    mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
        var results []SearchResult
        var err error

        results, err = search(r.FormValue("search"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        encoder := json.NewEncoder(w)
        err = encoder.Encode(results)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }).Methods("POST")






    mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request){
        var book ClassifyBookResponse
        var err error

        book, err = find(r.FormValue("id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        result, err := db.Exec("insert into books (pk, title, author, id, classification) values (?, ?, ?, ?, ?)",
                        nil, book.BookData.Title, book.BookData.Author, book.BookData.ID, book.Classification.MostPopular)
        if err != nil {
               http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        pk, _ := result.LastInsertId()
        b := Book{
        	PK: int(pk),
        	Title: book.BookData.Title,
        	Author: book.BookData.Title,
        	Classification: book.Classification.MostPopular,
        }

        if err := json.NewEncoder(w).Encode(b); err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }).Methods("PUT")





    mux.HandleFunc("/books/{pk}", func(w http.ResponseWriter, r *http.Request){
    	_, err := db.Exec("delete from books where pk = ?", gmux.Vars(r)["pk"])
        if err != nil {
               http.Error(w, err.Error(), http.StatusInternalServerError)
               return
        }
        w.WriteHeader(http.StatusOK)
    }).Methods("DELETE")

    n := negroni.Classic()
    n.Use(negroni.HandlerFunc(verifyDatabase))
    n.UseHandler(mux)
    n.Run(":8080")
}
