package main

import (
	"flag"
	"fmt"
	"encoding/pem"
	//"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

var (
	listenAddr string
)

type key int

const (
	requestIDKey key = 0
)

func main() {
	//flag.StringVar(&listenAddr, "listen-addr", ":4001", "server listen address")
	flag.Parse()

	//log.SetOutput(os.Stderr)
	logger := log.New(os.Stdout, "https", log.LstdFlags)

	router := http.NewServeMux()
	router.Handle("/", index())

	// nextRequestID := func() string {
	// 	return fmt.Sprintf("%d", time.Now().UnixNano())
	// }

	// server := &http.Server{
	// 	Addr:         listenAddr,
	// 	Handler:      router,
	// 	ErrorLog:     logger,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	//certPath := "./cert.pem"
	keyPath := "./key.pem"	

	//cert,_:= ioutil.ReadFile(certPath);
	key,_:= ioutil.ReadFile(keyPath);


	err := http.ListenAndServeTLS(":4001", "./cert.pem", key, nil)

	logger.Println(err)
	//server.ListenAndServe()
}

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, Tvrd!")
	})
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
