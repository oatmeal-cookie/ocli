package oatmealcookies

// export type AllRecipeJson = {
//     name: string;
//     recipeIngredient: string[];
//     recipeInstructions?: { text: string }[];
//     description: string;
//     image?: {
//       url: string;
//     };
//     video?: {
//       thumbnailUrl: string;
//     };
//     recipeYield: string[];
//   }[];

type AllRecipeJson struct {
	Name               *string   `json:"name"`
	RecipeIngredient   *[]string `json:"recipeIngredient"`
	RecipeInstructions *[]struct {
		Text *string `json:"text"`
	} `json:"recipeInstructions"`
	RecipeYeild *[]string `json:"recipeYield"`

	X map[string]interface{} `json:"-"` // catch all for unknown fields
}

func (aRJ AllRecipeJson) convertToFoundRecipe() FoundRecipe {

	instructions := make([]string, 0)
	instr := *aRJ.RecipeInstructions

	for i := 0; i < len(instr); i++ {
		str := instr[i].Text

		instructions = append(instructions, *str)
	}
	return FoundRecipe{
		Name:               aRJ.Name,
		RecipeIngredient:   aRJ.RecipeIngredient,
		RecipeInstructions: &instructions,
		RecipeYeild:        aRJ.RecipeYeild,
		X:                  aRJ.X,
	}
}

type FoundRecipe struct {
	Name               *string   `json:"name"`
	RecipeIngredient   *[]string `json:"recipeIngredient"`
	RecipeInstructions *[]string `json:"recipeInstructions"`
	RecipeYeild        *[]string `json:"recipeYield"`

	X map[string]interface{} `json:"-"` // catch all for unknown fields
}

type Recipe interface {
	convertToFoundRecipe() FoundRecipe
}
