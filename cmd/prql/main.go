package main

import (
	"github.com/mcandre/prql"

	"flag"
	"fmt"
	"os"
)

var flagVersion = flag.Bool("version", false, "Show version information")
var flagHelp = flag.Bool("help", false, "Show usage information")

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Println(prql.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	paths := flag.Args()

	for _, pth := range paths {
		if err := prql.Prql(pth); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
