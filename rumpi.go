package main

import (
	"github.com/FerdinaKusumah/rumpi/internal"
)

func main() {
	opt := internal.ParseOption()
	internal.RunWatch(opt)
}
