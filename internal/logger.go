package internal

import "context"

type LoggerIFace interface {
	Print(ctx context.Context, args ...interface{})
	Printf(ctx context.Context, format string, args ...interface{})
	Println(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Fatalln(ctx context.Context, args ...interface{})
	Panic(ctx context.Context, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Panicln(ctx context.Context, args ...interface{})
}
