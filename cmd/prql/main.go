package main

import (
	"github.com/mcandre/prql"

	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Usage: %s <path> [<path> [<path>...]]\n", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	paths := os.Args[1:]

	for _, pth := range paths {
		if err := prql.Prql(pth); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
