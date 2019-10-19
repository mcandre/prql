package prql

import (
	"github.com/xwb1989/sqlparser"

	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Version is semver.
const Version = "0.0.2"

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

	defer func() {
		if err3 := reader.Close(); err3 != nil {
			fmt.Println(err3)
		}
	}()

	tokenizer := sqlparser.NewTokenizer(reader)

	for {
		_, err3 := sqlparser.ParseNext(tokenizer)

		if err3 == io.EOF {
			break
		}

		if err3 != nil {
			var bytePosition int
			var line int
			var foundLineBreaks bool

			reader2, err2 := os.Open(pth)

			if err2 != nil {
				return err2
			}

			defer func() {
				if err3 := reader2.Close(); err3 != nil {
					fmt.Println(err3)
				}
			}()

			scanner := bufio.NewScanner(reader2)

			for scanner.Scan() {
				foundLineBreaks = true

				bytePosition += len(scanner.Text())

				if bytePosition > tokenizer.Position {
					break
				}

				line += 1
			}

			if !foundLineBreaks {
				line += 1
			}

			message := strings.Replace(err3.Error(), "syntax error at position", "syntax error at byte position", 1)

			return fmt.Errorf("%s:%d: %s", pth, line, message)
		}
	}

	return nil
}
