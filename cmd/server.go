package main

import (
	"math/rand"
	"time"

	"github.com/stingalleman/stingstore/server"
)

func main() {
	rand.Seed(time.Now().Unix())
	server.CreateServer("localhost", "9865")
}
