package main

import "Todo/microservice_auth/server"

func main() {


	s := server.NewServer(":3001")
	s.ListenAndServe()
}
