package main

import (
	"fmt"

	oatmealcookies "github.com/oatmeal-cookie/ocli/src/oatmealCookie"
)

func main() {
	recipe, err := oatmealcookies.UrlToJsonFile(
		"https://www.allrecipes.com/recipe/24202/shepherds-pie-vi/",
		"/Users/andrewalgard/recipejson/sheperds-pie.json")
	if err != nil {
		panic(err)
	}
	_, err = oatmealcookies.UrlToMarkdownFile(
		"https://www.allrecipes.com/recipe/24202/shepherds-pie-vi/",
		"/Users/andrewalgard/recipemarkdown/sheperds-pie.md")

	if err != nil {
		panic(err)
	}
	fmt.Println(recipe)

}
