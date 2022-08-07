package main

import (
	"golang-mazaya/storefront/internal/dialer"
	"log"
)

func main() {
	log.Println("Storefront running!")

	dialer.TestSave()

	dialer.Subscribe("recipe.created", func(msg []byte) {
		log.Printf("Received a message: %s", msg)
	})
}
