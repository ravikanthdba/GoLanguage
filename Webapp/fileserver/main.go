package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
)

func main() {
    r := httprouter.New();
    r.GET("/", HomeHandler);

    
    r.GET("/products", ShowProducts);
    r.POST("/products", CreateProducts);


    r.GET("/products/:id", ShowSpecificProduct);
    r.POST("/products/:id", CreateSpecificProduct);
    r.PUT("/products/:id", UpdateSpecificProduct);
    r.DELETE("/products/:id", DeleteSpecificProduct);

    port := os.Getenv("PORT")

    if port == "" {
        port = "8080";
    }

    fmt.Println("Starting WebApp on ", port)

    http.ListenAndServe(":"+port, r);
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Welcome to More Departmental Store and the value of p is : ", p);
	fmt.Fprintln(rw, "Welcome to More Departmental Store and the value of r is : ", r);
	fmt.Fprintln(rw, "Welcome to More Departmental Store and the value of rw is : ", rw);
}

func ShowProducts(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "This API shows the list of Products and the value of p is : ", p);
	fmt.Fprintln(rw, "This API shows the list of Products and the value of r is : ", r);
	fmt.Fprintln(rw, "This API shows the list of Products and the value of rw is : ", rw);
}

func CreateProducts(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "This API creates the list of Products and the value of p: ", p);
	fmt.Fprintln(rw, "This API creates the list of Products and the value of r: ", r);
	fmt.Fprintln(rw, "This API creates the list of Products and the value of rw: ", rw);
}

func ShowSpecificProduct(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id");
	fmt.Fprintln(rw, "This API shows a single product and the value of id is: ", id);
	fmt.Fprintln(rw, "This API shows a single product and the value of p is : ", p);
	fmt.Fprintln(rw, "This API shows a single product and the value of r is : ", r);
	fmt.Fprintln(rw, "This API shows a single product and the value of rw is : ", &rw);
} 

func CreateSpecificProduct(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
        id := p.ByName("id");
	fmt.Fprintln(rw, "Product " ,  id ,  " is created successfully");

}

func UpdateSpecificProduct(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
        id := p.ByName("id");
        fmt.Fprintln(rw, "Product " ,  id ,  "has successfully been updated");
}

func DeleteSpecificProduct(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
        id := p.ByName("id");
	fmt.Fprintln(rw, "Product "  , id ,  "has been deleted successfully");
}
