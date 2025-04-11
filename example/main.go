package main

import (
	"fmt"

	"github.com/a-random-lemurian/go-owot"
)

func main() {
	client, err := owot.Dial("wss://ourworldoftext.com/go-owot/ws/")
	if err != nil {
		panic(err)
	}

	client.HandleRaw = func(b []byte) {
		fmt.Printf("%s\n",b)
	}

	fmt.Println(client.Run())
}
