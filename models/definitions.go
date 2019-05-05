package models

const(
	MEAT_FOODS = 1
	FISH_FOODS = 2
	DAIRY_FOODS = 3
	EGG_FOODS = 4
	SPICES_AND_HERBS = 5
	BABY_FOODS =6
	FATS_AND_OILS = 7
	SOUPS_SAUCES_AND_GRAVIES =8
	BREAKFAST_CEREALS =9
	FRUIT =10
	VEGETABLE =11
	NUTS_AND_SEEDS =12
	BEVERAGES =13
	LEGUMES =14
	BAKED_PRODUCTS =15
	SWEETS =16
	CEREAL_GRAINS_AND_PASTA =18
	FAST_FOODS =19
	MEALS_ENTREES_AND_SIDE_DISHES =20
	SNACKS =21
	OTHER =22
	RESTAURANT_FOODS =23
)

const (
	RecipeCategoryNameVegetarian = "vegetarian"
	RecipeCategoryNameVegan = "vegan-friendly"
	RecipeCategoryNameGlutenFree = "gluten-free"
	RecipeCategoryNameLowFat = "low-fat"
	RecipeCategoryNameLowSatFat = "low-sat-fat"
	RecipeCategoryNameLowSugar = "low-sugar"
	RecipeCategoryNamePescatarian = "pescatarian"
	RecipeCategoryNameSavoryFood = "savory-food"
	RecipeCategoryNameSweetFood = "sweet-food"
)
var RecipeCategoryNames = [...]string{
	RecipeCategoryNameVegetarian,
	RecipeCategoryNameVegan,
	RecipeCategoryNameGlutenFree,
	RecipeCategoryNameLowFat,
	RecipeCategoryNameLowSatFat,
	RecipeCategoryNameLowSugar,
	RecipeCategoryNamePescatarian,
	RecipeCategoryNameSavoryFood,
	RecipeCategoryNameSweetFood,
}

const (
	RecipePreLoadIncludeCommonNames  = 1 << 1
	RecipePreLoadIncludeRecipePerPortion  = 1 << 2
	RecipePreLoadIncludeFoodNames  = 1 << 3
	RecipePreLoadIncludeFoodIds  = 1 << 4
	RecipePreLoadIncludeRecipeNutrition  = 1 << 5
	RecipePreLoadIncludeRecipeGDA  = 1 << 6
	RecipePreLoadIncludeCategories  = 1 << 7

	RecipePreLoadIncludeAll = RecipePreLoadIncludeCommonNames | RecipePreLoadIncludeRecipePerPortion  | RecipePreLoadIncludeFoodNames  | RecipePreLoadIncludeFoodIds |RecipePreLoadIncludeRecipeNutrition  |RecipePreLoadIncludeRecipeGDA | RecipePreLoadIncludeCategories
)

