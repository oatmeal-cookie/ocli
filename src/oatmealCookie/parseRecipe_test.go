package oatmealcookies

import (
	"testing"
)

type testArgs struct {
	url      string
	filepath string
}

func getTestArgs() []testArgs {
	ret := make([]testArgs, 0)

	ret = append(
		ret,
		testArgs{
			url:      "https://www.allrecipes.com/recipe/24202/shepherds-pie-vi/",
			filepath: "/Users/andrewalgard/recipemarkdown/",
		})
	return ret
}

func TestCreateMarkdown(t *testing.T) {
	toTest := getTestArgs()

	for i := 0; i < len(toTest); i++ {
		testCase := toTest[i]
		_, err := UrlToMarkdownFile(
			testCase.url,
			testCase.filepath,
		)

		if err != nil {
			t.Fatal(err.Error())
			break
		}
	}
}
