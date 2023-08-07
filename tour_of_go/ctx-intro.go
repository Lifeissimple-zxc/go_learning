package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	// This WithValue adds a {"request-id": "12345"} kv pair
	// To the context struct
	return context.WithValue(ctx, "request-id", "12345")
}

func doSmthCool(ctx context.Context) {
	rID := ctx.Value("request-id") // Value retrieves request-id key from the ctx struct
	fmt.Println(rID)
}

// doSmthElse shows how context.Context.Done() works
func doSmthElse(ctx context.Context) {
	// rID := ctx.Value("request-id") // Value retrieves request-id key from the ctx struct
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timed out")
			return
		default:
			fmt.Println("Doing smth else")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting CTX demo")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	fmt.Println("doSmthCool:")
	doSmthCool(ctx)
	fmt.Println("Starting Timed Context Demo")
	newCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // This closes our context
	// This is a bit tricky
	go doSmthElse(newCtx) // launch doSmthelse on a separate thread
	select {
	case <-newCtx.Done():
		fmt.Println("Deadline exceeded")
	}

	time.Sleep(2 * time.Second)

}
