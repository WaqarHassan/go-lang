package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	mySleepAndTalk(ctx, 2*time.Second, "hello")
}

func mySleepAndTalk(ctx context.Context, t time.Duration, msg string) {
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
