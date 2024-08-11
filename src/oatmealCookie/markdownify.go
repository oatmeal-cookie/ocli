package oatmealcookies

import (
	"os"
	"strings"

	md "github.com/nao1215/markdown"
)

func generateMarkdown(data FoundRecipe, filepath string) error {
	ingredients := *data.RecipeIngredient
	instructions := *data.RecipeInstructions
	servings := *data.RecipeYeild

	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	builder := md.NewMarkdown(f).
		H1(*data.Name)
	if len(servings) > 0 {
		builder.PlainText("Serves " + servings[0])
	}
	builder.PlainText("\n")
	builder.H2("Ingredients")
	var ingredientsBuilder strings.Builder
	for i := 0; i < len(ingredients); i++ {
		ingredientsBuilder.WriteString(ingredients[i] + "  \n")
	}
	builder.PlainText(ingredientsBuilder.String())
	builder.
		PlainText("\n").
		H2("Instructions").
		BulletList(instructions...).
		PlainText("\n")

	return builder.Build()
}
