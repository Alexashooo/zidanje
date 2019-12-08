package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	listenAddr string
)

type key int

const (
	requestIDKey key = 0
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":4001", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "https", log.LstdFlags)
	http.Handle("/", http.FileServer(http.Dir("./ClientApp/dist/ClientApp")))

	err := http.ListenAndServeTLS(listenAddr, "cert.pem", "key.unencrypted.pem", nil)

	logger.Println(err)
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
