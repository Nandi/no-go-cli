package main

import _ "embed"

import (
	"fmt"
	"math/rand"
	"strings"
)

//go:embed reasons.txt
var reasons string

func main() {
	payload := strings.Split(reasons, "\n")

	index := rand.Intn(len(payload))

	fmt.Println(payload[index])
}
