package main

import (
	"context"
	"fmt"
	"time"

	"github.com/a-random-lemurian/go-owot"
)

func main() {
	client, err := owot.Dial("wss://ourworldoftext.com/go-owot/ws/", nil)
	if err != nil {
		panic(err)
	}

	client.HandleRaw = func(b []byte) {
		fmt.Printf("%s\n", b)
	}

	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go client.RunContext(ctx)

	time.Sleep(5 * time.Second)

	cancel()
}
