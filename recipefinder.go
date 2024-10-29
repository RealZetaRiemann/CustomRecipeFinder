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

var recipes []Recipe

func main() {
	loadRecipes("recipes.json")
	//fmt.Print(recipes[0])

	fmt.Print("\nThis program allows you to search for recipes that contain specific ingredients.\n\n")
	fmt.Print("You may enter as many ingredients as you like and the program will return the five recipes that are the best match.\n\n")
	fmt.Print("All ingredient names entered should be singular.\n\n")
	userIngredients := getUserIngredients()
	fmt.Print(userIngredients)
}

func loadRecipes(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening recipes file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&recipes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func getUserIngredients() []string {
	var Ingredients []string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many items would you like to input? ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // remove newline character

	// try to convert input to int
	numIngredients, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return Ingredients
	}

	for i := 0; i < numIngredients; i++ {
		fmt.Printf("Enter item %d: ", i+1)
		ingredient, _ := reader.ReadString('\n')
		ingredient = ingredient[:len(ingredient)-1] // remove newline character
		Ingredients = append(Ingredients, strings.ToLower(ingredient))
	}
	return Ingredients
}
