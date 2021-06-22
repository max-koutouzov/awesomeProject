package main

import (
	"context"
	"os"

	. "server/client"
	. "server/serverContext"
)

func main() {
	ss := SlowServer()
	defer ss.Close()
	fs := FastServer()
	defer fs.Close()

	ctx := context.Background()
	CallBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
