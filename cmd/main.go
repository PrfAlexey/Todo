package main

import (
	"Todo/server"
)

func main() {

	srv := server.NewServer()
	srv.ListenAndServe()
}
