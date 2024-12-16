// go run recipefinder.go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Recipe defines the information available for each recipe
type Recipe struct {
	Cookbook      string
	Recipe        string
	Page          int
	Ingredients   []string // might be better to make this a map and the user entered ingredients a slice?
	RoniaApproved bool
}

// Match is used for finding the indices of the recipes with the highest number of ingredient matches
type Match struct {
	NumMatches  int
	RecipeIndex int
}

var recipes []Recipe

func main() {
	loadRecipes("recipes.json")
	//fmt.Print(len(recipes))
	//fmt.Print(recipes[0])

	fmt.Print("\nThis program allows you to search for recipes that contain specific ingredients.\n\n")
	fmt.Print("You can choose how many ingredients to enter and how many recipes you'd like to recieve.\n\n")
	fmt.Print("The program will return the recipes that are the best match (include the most ingredients you requested).\n\n")
	fmt.Print("All ingredient names entered should be singular (e.g., mushroom instead of mushrooms).\n\n")
	userIngredients, numRecipes := getUserIngredients() // get user input
	//fmt.Print(userIngredients)
	//fmt.Print(numRecipes)
	//fmt.Print(findBestMatches(userIngredients, numRecipes))
	BestMatches := findBestMatches(userIngredients, numRecipes)
	printBestMatches(BestMatches)
}

func loadRecipes(filename string) {
	// open json file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening recipes file:", err)
		return
	}
	defer file.Close()

	// convert json file to objects
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&recipes) // decoding to recipes (list of Recipe objects)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func getUserIngredients() (map[string]int, int) {
	Ingredients := make(map[string]int) // initialize map of user-entered ingredients (more efficient than slice?)

	// get number of ingredients to ask for
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many ingredients would you like to input? ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // remove newline character

	// try to convert input (number of ingredients) to int
	numIngredients, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return Ingredients, 0
	}

	// continue to ask for ingredients depending on how many ingredients the user said they'd like to enter
	for i := 0; i < numIngredients; i++ {
		fmt.Printf("Enter ingredient %d: ", i+1)
		ingredient, _ := reader.ReadString('\n')
		ingredient = ingredient[:len(ingredient)-1] // remove newline character
		ingredient = strings.ToLower(ingredient)    //convert to lowercase (all lowercase in json)
		// if the user input a specific mushroom, also include the "generic" mushroom
		if MushroomClause(ingredient) {
			Ingredients["mushroom"] = 0
		}
		Ingredients[ingredient] = 0 // add ingredient to map
	}

	// get number of "best-match" recipes to find and return
	fmt.Print("How many recipes would you like to recieve? ")
	input, _ = reader.ReadString('\n')
	input = input[:len(input)-1] // remove newline character
	// try to convert input to int
	numRecipes, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return Ingredients, 0
	}

	return Ingredients, numRecipes
}

func findBestMatches(userIngredients map[string]int, numRecipes int) []Match {
	BestMatches := make([]Match, numRecipes) // make the list which with the indices of the best recipes with the most matches

	for r := 0; r < len(recipes); r++ { // for each recipe...
		numMatches := 0
		for i := 0; i < len(recipes[r].Ingredients); i++ { // for each ingredient...
			_, ok := userIngredients[recipes[r].Ingredients[i]] // if the ingredient is one of those requested by the user...
			if ok {
				numMatches++ // then increase the current number of matches for this recipe
			}
		}
		for i := 0; i < numRecipes; i++ { // for each best match recipe so far
			if numMatches > BestMatches[i].NumMatches { // if the new recipe has more ingredient matches than any of the current best matches
				BestMatches[i].NumMatches = numMatches // update the number of matches
				BestMatches[i].RecipeIndex = r         // update the recipe index
				break                                  // only need to replace one "best match" recipe
			}
		}
	}
	return BestMatches
}

func printBestMatches(BestMatches []Match) {
	fmt.Print("\n")
	for m := 0; m < len(BestMatches); m++ { // for each best match...
		recipe := recipes[BestMatches[m].RecipeIndex] // get the recipe object
		fmt.Printf("Recipe: %s\n", recipe.Recipe)
		fmt.Printf("Cookbook: %s\n", recipe.Cookbook)
		fmt.Printf("Page number: %d\n", recipe.Page)
		fmt.Printf("Number of ingredient matches: %d\n", BestMatches[m].NumMatches)
		fmt.Print("Ingredients: ")
		for i := 0; i < len(recipe.Ingredients); i++ {
			fmt.Printf("%s", recipe.Ingredients[i])
			if i != (len(recipe.Ingredients) - 1) {
				fmt.Print(",")
			}
			fmt.Print(" ")
		}
		fmt.Printf("\nRonia approved?: %t\n", recipe.RoniaApproved)
		fmt.Print("\n")
	}
}

// MushroomClause checks if the user input a specific type of mushroom...
// then the function should also search for the "generic" mushroom
func MushroomClause(userInput string) bool {
	if strings.Contains(userInput, "mushroom") && userInput != "mushroom" {
		return true
	}
	return false
}
