package main

import (
	//"os"
	"fmt"
	"time"
	"math/rand"

	"woctl/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewWoyectl()
	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}
