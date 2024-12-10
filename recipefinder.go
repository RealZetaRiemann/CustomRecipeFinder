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
	Ingredients   []string
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
	//fmt.Print(recipes[0])

	fmt.Print("\nThis program allows you to search for recipes that contain specific ingredients.\n\n")
	fmt.Print("You can choose how many ingredients to enter and how many recipes you'd like to recieve.\n\n")
	fmt.Print("The program will return the recipes that are the best match (include the most ingredients you requested).\n\n")
	fmt.Print("All ingredient names entered should be singular.\n\n")
	userIngredients, numRecipes := getUserIngredients() // get user input
	fmt.Print(userIngredients)
	fmt.Print(numRecipes)
	//findBestMatches(userIngredients,numRecipes)
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
		ingredient = ingredient[:len(ingredient)-1]  // remove newline character
		Ingredients[strings.ToLower(ingredient)] = 0 // convert all ingredients to lowercase (all lowercase in json)
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

func findBestMatches(userIngredients map[string]int, numRecipes int) {
	BestIndices := make([]Match, numRecipes) // make the list which with the indices of the best recipes with the most matches

	for r := 0; r < len(recipes); r++ { // for each recipe...
		for i := 0; i < len(recipes[r].Ingredients); i++ { // for each ingredient...
			numMatches := 0
			_, ok := userIngredients[recipes[r].Ingredients[i]] // if the ingredient is one of those requested by the user...
			if ok {
				numMatches++ // then increase the current number of matches for this recipe
			}
		}
		// TBD check if numMatches > the values in BestIndices & if it is, then replace the highest value that it is greater than
	}
}
