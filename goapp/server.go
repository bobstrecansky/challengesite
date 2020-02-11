package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type server struct {
	router *http.ServeMux
	db     *sql.DB
	cache  *redis.Client
}

func challengeAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Challenge App!")
}

func redisPingHandler(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, pong)
}

// These need to go to ENV variables, just putting them here for now
const (
	host     = "postgres"
	port     = "5432"
	user     = "postgres"
	password = "test123"
	dbname   = "challenge"
)

func postgresPingHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "host=postgres user=postgres dbname=challenge password=test123 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, "Successfully connected to postgres")
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", challengeAppHandler)
	r.HandleFunc("/redisPing", redisPingHandler)
	r.HandleFunc("/postgresPing", postgresPingHandler)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
