package oatmealcookies

import (
	"bufio"
	"bytes"
	"testing"
)

func getTestArgs() []string {

	return []string{
		"https://www.allrecipes.com/recipe/24202/shepherds-pie-vi/",
	}
}

func TestCreateMarkdown(t *testing.T) {
	toTest := getTestArgs()

	for i := range len(toTest) {
		testCase := toTest[i]
		var b bytes.Buffer
		writer := bufio.NewWriter(&b)
		_, err := UrlToMarkdown(
			testCase,
			writer,
		)
		t.Log(b.String())

		if err != nil {
			t.Fatal(err.Error())
			break
		}

	}
}
