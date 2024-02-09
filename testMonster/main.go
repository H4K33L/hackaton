package main

import (
	"Monster"
	"fmt"
)

func main() {
	err := Monster.Generate("Aberration", 3)
	if err != nil {
		fmt.Println(err)
	}
}
