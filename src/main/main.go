package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	oatmealcookies "github.com/oatmeal-cookie/ocli/src/oatmealCookie"
)

var format string
var filepath string

func init() {

	flag.StringVar(&format, "format", "md", "output format of recipe")
	flag.StringVar(&format, "f", "", "output format of recipe")
}

func firstFlag() int {
	flagIndex := 1
	for i := 1; i < len(os.Args); i++ {
		flagIndex = i
		if strings.HasPrefix(os.Args[i], "-") {
			break
		}
	}
	return flagIndex
}

func main() {
	flagIndex := firstFlag()
	flag.CommandLine.Parse(os.Args[flagIndex:])

	args := os.Args
	if len(args) < 2 {
		os.Stderr.WriteString("Missing arg: URL is required")
		return
	}
	url := args[1]
	parseUrl(format, url, filepath)
}

func parseUrl(format string, url string, filepath string) {
	writer, err := createOutput(filepath)
	if v, ok := writer.(*os.File); ok {
		defer v.Close()
	}

	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	if format == "md" {
		_, err := oatmealcookies.UrlToMarkdown(
			url,
			writer)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
		return
	}
	if format == "json" {
		_, err := oatmealcookies.UrlToJsonFile(
			url,
			writer)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
	}
}

func createOutput(filepath string) (io.Writer, error) {
	if filepath == "" {
		return os.Stdout, nil
	} else {
		valid := IsValid(filepath)
		if valid {
			file, err := os.Create(filepath)
			if err != nil {
				return file, err
			}
			return file, nil
		} else {
			return nil, errors.New("invalid filepath")
		}
	}

}

func IsValid(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := os.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}
