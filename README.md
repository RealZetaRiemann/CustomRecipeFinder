This program was made as a gift for a family member to search through their personal cookbook collection and find recipes that contain specific ingredients.

recipefinder.go will ask the user how many ingredients they would like to search for. It will then prompt the user to enter each of those ingredients one-by-one. The user will then asked how many recipes they would like to recieve. 

Based on this information, recipefinder.go will print information about the recipes with the most ingredient matches. 

The information printed includes the:
- Recipe (name of recipe)
- Cookbook (name of cookbook)
- Page Number
- Number of "Ingredient Matches" (number of user-inputted ingredients it uses)
- Ingredients (all ingredients in the recipe)
- Ronia Approval (whether or not I personally would like to try eating it)

When entering ingredients, please make sure to follow these rules to get the best-possible results:
- All ingredients should always be plural, even if it sounds weird (e.g., peas -> pea, mushrooms -> mushroom)
- Try not to be too specific (e.g., garlic clove -> garlic, light brown sugar -> brown sugar, firm silken tofu -> tofu)
    - This rule really depends on the ingredients, if you aren't sure you can look in recipes.json to see what I did

Misc. Notes:
- Water, salt, and black pepper are not included as ingredients for the sake of brevity
- Optional ingredients are not included
- Mushrooms *are* listed by the specific type (i.e., shiitake mushroom)
- Caster sugar, white sugar, granulated sugar, etc. are all just sugar
- Silken tofu, firm tofu, etc. are all just tofu
- Brown/yellow onions are both just onion, but red onion and scallion are seperate
- All recipes are vegetarian, which means any meat is just assumed to be vegan
- Cherry tomatoes, grape tomatoes, etc. are typically all listed as tomato
- If a recipe calls for a generic oil it is not included, but if a specific oil is listed in the recipe, then it is included
- Mayonnaise is shortened to mayo
- Brown rice, jasmine rice, etc. are all just rice

Feedback and ideas are always welcome!