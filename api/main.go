package main

import (
	"auth-server/api/handlejwt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type check struct {
	auth bool
}

func main() {
	r := mux.NewRouter()
	r.Handle("/api/v1/auth/check", handlejwt.JwtMiddleware.Handler(checkauth))
	r.Handle("/api/v1/auth/gettoken", handlejwt.GetTokenHandler)

	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var checkauth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	c := &check{
		auth: true,
	}
	json.NewEncoder(w).Encode(c)
})
