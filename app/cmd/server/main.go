package main

import (
	"log"

	checkouts "github.com/jeancarlosdanese/go-tracing/checkouts/cmd/server"
	payments "github.com/jeancarlosdanese/go-tracing/payments/cmd/server"
)

func main() {
	log.Println("running applications")

	go func() {
		checkouts.Serve()
	}()

	payments.Serve()
}
