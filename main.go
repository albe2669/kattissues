package main

import (
	"github.com/albe2669/kattissues/cmd"
	"github.com/albe2669/kattissues/internal"
)

func main() {
	err := internal.ReadConfig()
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
