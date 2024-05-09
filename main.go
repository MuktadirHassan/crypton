package main

import "github.com/MuktadirHassan/crypton/api"

func main() {
	server := api.NewServer(":8081")

	err := server.Start()
	if err != nil {
		panic(err)
	}

}
