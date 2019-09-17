package main

import (
	"math/rand"
	"time"

	"github.com/stringUtil/cmd/stringUtil/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	cmd.Execute()
}
