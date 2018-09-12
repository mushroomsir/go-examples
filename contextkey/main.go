package main

import (
	"context"
	"fmt"
)

type contextKey struct{}

var activeSpanKey = contextKey{}

var activeSpanKey2 = contextKey{}

func main() {
	ctx := context.WithValue(context.Background(), activeSpanKey, "cc")
	fmt.Println(ctx.Value(activeSpanKey2))
	fmt.Println(activeSpanKey == activeSpanKey2)
}
