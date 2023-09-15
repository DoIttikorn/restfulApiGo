package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Age  int
	Name string
}

var users = []User{}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		method := r.Method

		if method == "GET" {

			b, err := json.Marshal(users)
			if err != nil {
				w.Write([]byte("json marshal error"))
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))

			return
		}

		if method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("read body error"))
			}
			var u = &User{}

			err = json.Unmarshal(body, u)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("json unmarshal error"))
			}
			u.Id = len(users) + 1

			users = append(users, *u)

			w.Write([]byte("success"))
			return
		}

		w.Write([]byte("not support method"))
	})

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
