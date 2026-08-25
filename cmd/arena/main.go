package main

import (
	"example.com/emblem/api"
	"example.com/emblem/service"
	"example.com/emblem/storage"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	path := flag.String("db", "arena.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	s, e := storage.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	srv := api.New(service.New(s))
	if e := http.ListenAndServe(*addr, srv.Routes()); e != nil && !os.IsTimeout(e) {
		log.Fatal(e)
	}
}
