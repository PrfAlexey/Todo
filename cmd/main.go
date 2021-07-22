package main

import (
	"ads/server"
)

func main() {

	srv := server.NewServer()
	srv.ListenAndServe()
}
