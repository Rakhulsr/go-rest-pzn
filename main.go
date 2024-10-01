package main

import (
	"net/http"

	"github.com/Rakhulsr/go-rest-api-pzn/helper"
	"github.com/Rakhulsr/go-rest-api-pzn/middleware"
)

func NewServer(auth *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: auth,
	}
}

func main() {
	server := InitServer()

	// log.Println("Listening to", "http://"+server.Addr)
	err := server.ListenAndServe()
	helper.PanicError(err)

}

// func initDb(db *sql.DB) {
// 	err := db.Ping()
// 	helper.PanicError(err)

// 	log.Println("DB CONNECTED!")
// }
