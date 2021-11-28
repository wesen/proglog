package main

import (
	"github.com/wesen/proglog/internal/server"
	"log"
)

//go:generate protoc api/v1/log.proto --go_out=../.. --go_opt=paths=source_relative --proto_path=../..

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
