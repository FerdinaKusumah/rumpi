package main

import (
	"example/watch-api/internal"
)

func main() {
	opt := internal.ParseOption()
	internal.RunWatch(opt)
}
