package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	// test
	fmt.Print(recipes[0].Page)
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
