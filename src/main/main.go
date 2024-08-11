package main

import (
	"flag"
	"os"
	"strings"

	oatmealcookies "github.com/oatmeal-cookie/ocli/src/oatmealCookie"
)

var format string

func init() {
	flag.StringVar(&format, "format", "md", "output format of recipe")
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
		os.Stderr.WriteString("Missing args: URL and output filepath required")
		return
	}
	if len(args) < 3 {
		os.Stderr.WriteString("Missing arg: output filepath required")
		return
	}
	url := args[1]
	filepath := args[2]
	parseUrl(format, url, filepath)
}
func parseUrl(format string, url string, filepath string) {
	if format == "md" {
		_, err := oatmealcookies.UrlToMarkdownFile(
			url,
			filepath)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
		return
	}
	if format == "json" {

		_, err := oatmealcookies.UrlToJsonFile(
			url,
			filepath)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
	}
}
