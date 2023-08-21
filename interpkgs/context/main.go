package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "my value")
	ctx = context.WithValue(ctx, "my-key-int", 5)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go myProcess(ctx2)

	select {
	case <- ctx2.Done():
		fmt.Println("we've exceeded the deadline")
		fmt.Println(ctx2.Err())
	}

}

func viewContext(ctx context.Context) {
	fmt.Printf("my value is '%s'\n", ctx.Value("my-key"))
	fmt.Printf("my other value is %d\n", ctx.Value("my-key-int"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ohh, time out")
			return
		default:
			fmt.Println("doing some process")
		}
		time.Sleep(500 * time.Millisecond)
	}
}