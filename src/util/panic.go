package util

import (
	"context"
	"log"
)

func PanicHandle(f func(context.Context)) func(context.Context) {
	return func(ctx context.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				log.Fatalf("panic: %v", msg)
			}
		}()
		f(ctx)
	}
}
