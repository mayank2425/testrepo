package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	// "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


func main() {
	router := httprouter.New()
	router.GET("/",index)
    log.Fatal(http.ListenAndServe(":8080",router))
}