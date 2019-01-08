package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, world.", "My favorite number is", rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Pi)
}
