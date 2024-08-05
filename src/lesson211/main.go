package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func Set(userID, authToken string) context.Context {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	return ctx
}

func Get(ctx context.Context) {
	fmt.Println(ctx.Value(ctxUserID).(string))
	fmt.Println(ctx.Value(ctxAuthToken).(string))
}

func main() {
	ctx := Set("123", "abc")
	Get(ctx)
}
