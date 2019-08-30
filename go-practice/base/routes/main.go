package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"

	_ "github.com/drone/routes"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are get user %s", uid)
}

func modifyUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are modify user %s", uid)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are delete user %s", uid)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are add user %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getUser)
	mux.Put("/user/:uid", modifyUser)
	mux.Del("/user/:uid", deleteUser)
	mux.Post("/user/:uid", addUser)

	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
