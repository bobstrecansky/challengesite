package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type server struct {
	router *http.ServeMux
	db     *sql.DB
	cache  *redis.Client
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func redisPingHandler(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, pong)
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", yourHandler)
	r.HandleFunc("/redisPing", redisPingHandler)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
