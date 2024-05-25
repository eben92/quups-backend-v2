package main

import (
	"fmt"

	"quups-backend/internal/server"
)

func main() {
	server := server.NewServer()
	fmt.Print("rr")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
