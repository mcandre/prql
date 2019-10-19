package main

import (
	"github.com/xwb1989/sqlparser"

	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func usage() {
	fmt.Printf("Usage: %s <path> [<path> [<path>...]]\n", os.Args[0])
}

func prql(pth string) error {
	fi, err := os.Stat(pth)

	if err != nil {
		return err
	}

	if fi.Mode().IsDir() {
		childInfos, err2 := ioutil.ReadDir(pth)

		if err2 != nil {
			return err2
		}

		for _, childInfo := range childInfos {
			if err3 := prql(path.Join(pth, childInfo.Name())); err3 != nil {
				return err3
			}
		}

		return nil
	}

	reader, err2 := os.Open(pth)

	if err2 != nil {
		return err2
	}

	tokenizer := sqlparser.NewTokenizer(reader)

	for {
		_, err3 := sqlparser.ParseNext(tokenizer)

		if err3 == io.EOF {
			break
		}

		if err3 != nil {
			return fmt.Errorf("%s: %v", pth, err3)
		}
	}

	return reader.Close()
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	paths := os.Args[1:]

	for _, pth := range paths {
		if err := prql(pth); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
