package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func hello() {
	fmt.Println("Hello Go Usergroup Rhein-Ruhr")
}

func alternatingHello() {
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return
	}
	if n.Int64() == 0 {
		fmt.Println("Hello World")
		fmt.Println("Hello Go Usergroup Rhein-Ruhr")
	} else {
		fmt.Println("Hello Go Usergroup Rhein-Ruhr")
		fmt.Println("Hello World")
	}
}

func main() {
	hello()
}
