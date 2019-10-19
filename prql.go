package prql

import (
	"github.com/xwb1989/sqlparser"

	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Version is semver.
const Version = "0.0.1"

// Prql recursively scans a directory or file path for SQL scripts.
// On successful parsing of all files, returns nil.
// Otherwise, returns a parse error.
func Prql(pth string) error {
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
			if err3 := Prql(path.Join(pth, childInfo.Name())); err3 != nil {
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
