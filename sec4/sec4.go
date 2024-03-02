package main

import (
	"fmt"
)

func main() {
	words := []string{"hi", "salutuations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "short", wordLen)
		case wordLen > 10:
			fmt.Println(word, "long", wordLen)
		}
	}
}
