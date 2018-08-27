package main

import (
	"context"
	"time"

	context1 "golang.org/x/net/context"
)

// ContextX ...
type ContextX = context1.Context

func main() {
	context.WithTimeout(context1.Background(), 30*time.Second)
}
