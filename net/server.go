package net

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ServerLocal() {

	router := mux.NewRouter()

	router.HandleFunc("/", Inicial)
	router.HandleFunc("/newLink", PostLink).Methods(http.MethodPost)

	http.ListenAndServe(":8081", router)

}

/*

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi everyone, i like %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

*/
