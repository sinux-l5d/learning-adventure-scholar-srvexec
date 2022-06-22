package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	// if an arg exists, use it as the input
	if len(os.Args) > 1 {
		input := os.Args[1]
		decoded, _ := base64.StdEncoding.DecodeString(input)
		fmt.Printf("%s\n", decoded)
	} else {
		// While not SIGINT, scan for input
		for {
			var input string
			fmt.Scanln(&input)
			decoded, err := base64.StdEncoding.DecodeString(input)
			if err == nil {
				fmt.Printf("%s\n", decoded)
			}
		}
	}
}
