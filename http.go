package main

import (
	crdt "github.com/pingles/crdt-go"
	"log"
	"net/http"
	"strconv"
)

func writeCounter(counter *crdt.Counter, w http.ResponseWriter) {
	w.Write([]byte(strconv.FormatInt(counter.Value(), 10)))
}

func ServeCounter(counter *crdt.Counter, httpBinding string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeCounter(counter, w)
	})
	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		counter.Increment()
		writeCounter(counter, w)
	})
	http.HandleFunc("/dec", func(w http.ResponseWriter, r *http.Request) {
		counter.Decrement()
		writeCounter(counter, w)
	})

	log.Println("http listening", httpBinding)
	http.ListenAndServe(httpBinding, nil)
}
