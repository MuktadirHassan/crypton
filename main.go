package main

import (
	"fmt"

	"github.com/MuktadirHassan/crypton/pkg/rsa"
)

func main() {
	// server := api.NewServer(":8081")

	// err := server.Start()
	// if err != nil {
	// 	panic(err)
	// }

	p := rsa.GeneratePrimeNumber(100)
	q := rsa.GeneratePrimeNumber(100)
	fmt.Println(p)
	fmt.Println(q)
}
