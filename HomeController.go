package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
)

var (
	listenAddr string
)

type key int

const (
	requestIDKey key = 0
)


func BuildDB() bool{
	connStr := "user=gotest password=gotest dbname=postgres"
	
	db, err := sql.Open("postgres", connStr)

	if err !=nil {
		log.Fatal(err)
		return false;
	}
	
	test, _ := db.Query("SELECT 1")

	if test != nil {
		return true
	}

	_,err3 := db.Query("CREATE DATABASE testiclebaza") 

	if err3 !=nil {
		log.Fatal(err3)
		return false;
	}

	return true
}

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":4001", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "https", log.LstdFlags)

	if BuildDB() {
		http.Handle("/", http.FileServer(http.Dir("./ClientApp/dist/ClientApp/")))
		err := http.ListenAndServeTLS(listenAddr, "cert.pem", "key.unencrypted.pem", nil)
		logger.Println(err)
	}
}

// func logging(logger *log.Logger) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			defer func() {
// 				requestID, ok := r.Context().Value(requestIDKey).(string)
// 				if !ok {
// 					requestID = "unknown"
// 				}
// 				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
// 			}()
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			requestID := r.Header.Get("X-Request-Id")
// 			if requestID == "" {
// 				requestID = nextRequestID()
// 			}
// 			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
// 			w.Header().Set("X-Request-Id", requestID)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }
