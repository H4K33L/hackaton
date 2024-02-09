package main

import (
	"Monster"
	"fmt"
)

func main() {
	err := Monster.Generate(1, "Aberration", 3)
	if err != nil {
		fmt.Println(err)
	}
}

