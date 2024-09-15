// server
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"io"
	"strings"
)

var database Database

func PostHandler(response http.ResponseWriter, request *http.Request) {
	var buffer strings.Builder
	io.Copy(&buffer, request.Body)
	value := buffer.String()

	fmt.Println("POST: ", value)

	key := database.Insert(value)
	fmt.Fprint(response, key)
	
	fmt.Println("Response: ", key)
}

func GetHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]

	fmt.Println("GET: ", key)

	value, ok := database.GetByKey(key)

	if !ok {
		value = ""
	}
	fmt.Fprint(response, value)
	
	fmt.Println("Response: ", value)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", PostHandler).Methods("POST")
	router.HandleFunc("/{key:[a-zA-Z0-9]+}", GetHandler).Methods("GET")

	//parse
	useDatabase := flag.Bool("d", false, "Use a database? Default is false")
	flag.Parse()

	//show flags
	fmt.Println("UseDatabase: ", *useDatabase)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
	})

	if *useDatabase {
		connStr := "user=postgres password=password dbname=sitesdb sslmode=disable"
		db_, err := sql.Open("postgres", connStr)

		if err != nil {
			panic(err)
		}

		defer db_.Close()

		database = &PostgresDatabase{db: db_}

		db_.Exec(`CREATE TABLE sites
						(
						Id SERIAL PRIMARY KEY,
		    			Key CHARACTER VARYING(10) UNIQUE,
		    			Value CHARACTER VARYING(512) UNIQUE
						);`)
	} else {
		database = &MyDatabase{mKeyToValue: make(map[KeyType]ValueType), mValueToKey: make(map[ValueType]KeyType), top: 0}
	}
	fmt.Println("Server started")
	http.ListenAndServe("localhost:8080", router)
	fmt.Println("Server closed")
}
