package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ara-ta3/go-message-app-exercise/app"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		m := app.Message{
			ID:      "1",
			Message: "pong",
		}

		res, _ := json.Marshal(m)
		fmt.Fprintf(w, string(res))
	})
	fmt.Println("http server started on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// func withJsonContentTypeHeader(handleFunc func(w http.ResponseWriter, r *http.Request)) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "application/json")
//         handleFunc(w, r)
//     })
// }
