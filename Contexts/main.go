package main

import (
	"context"
	"fmt"
	"time"
)

// Simulating Data Fetching
func data(ctx context.Context) error {
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("Data Fetching!!!")
	case <-ctx.Done():
		return ctx.Err() // if request gets cancelled or timed out due to some reason
	}
	return nil
}

func main() {

	//	Contexts are used to maintain the lifecycle of requests for concurrent applications in Golang.
	//	Core Features:
	//	1. Cancel Propagation: Cancel a request when another request is done
	//	2. Deadlines: Set a deadline for an operation to prevent long-running tasks.
	//	3. Request-scoped Values: Pass values like authentication tokens or user IDs across function calls.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	// will give error since context states the operation should be performed in 1 second and the function takes 5
	defer cancel()

	if err := data(ctx); err != nil {
		fmt.Println("Error:", err)
	}
}
