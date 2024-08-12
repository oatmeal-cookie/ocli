package oatmealcookies

import (
	"encoding/json"
	"errors"
	"html"
	"io"
	"strings"

	"github.com/gocolly/colly"
)

func ExtractLdJson(recipeUrl string) (FoundRecipe, error) {
	c := colly.NewCollector()
	pJsonMap := new([]AllRecipeJson)
	// Find and visit all links
	c.OnHTML("script[type=\"application/ld+json\"]", func(e *colly.HTMLElement) {
		unescaped := html.UnescapeString(e.Text)
		json.Unmarshal([]byte(unescaped), &pJsonMap)
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("extracing html", r.URL)
	// })

	c.Visit(recipeUrl)
	if len(*pJsonMap) > 0 {
		first := (*pJsonMap)[0]
		return first.convertToFoundRecipe(), nil
	}
	return FoundRecipe{}, errors.New("no ld+json found")
}

func UrlToJsonFile(recipeUrl string, writer io.Writer) (FoundRecipe, error) {
	recipe, err := ExtractLdJson(recipeUrl)
	if err != nil {
		return FoundRecipe{}, err
	}
	// attempt to write to supplied file path
	// if error, return error
	// if success, return recipe
	toWrite, err := json.Marshal(recipe)
	if err != nil {
		return FoundRecipe{}, err
	}
	err = writeDataToFile(toWrite, writer)
	if err != nil {
		return FoundRecipe{}, err
	}
	return recipe, nil
}

func writeDataToFile(data []byte, writer io.Writer) error {
	_, err := writer.Write(data)
	return err
}

func getFilenameOfRecipe(recipe FoundRecipe, filepath string) string {
	var sb strings.Builder
	sb.WriteString(filepath)
	sb.WriteString(*recipe.Name)
	sb.WriteString(".md")
	return sb.String()
}

func UrlToMarkdown(recipeUrl string, output io.Writer) (FoundRecipe, error) {
	recipe, err := ExtractLdJson(recipeUrl)
	if err != nil {
		return FoundRecipe{}, err
	}
	err = generateMarkdown(recipe, output)
	return recipe, err
}
