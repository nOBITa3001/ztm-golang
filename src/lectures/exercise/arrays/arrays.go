//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printShoppingListInfo(list [4]Product) {
	var cost, totalItem int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price
		if item.name != "" {
			totalItem += 1
		}
	}

	fmt.Println("The last item on the list:", list[totalItem-1])
	fmt.Println("The total number of items:", totalItem)
	fmt.Println("The total cost of the items:", cost)
}

func main() {
	shoppingList := [4]Product{
		{name: "Banana", price: 1},
		{name: "Meat", price: 6},
		{name: "Salad", price: 3},
	}

	printShoppingListInfo(shoppingList)

	shoppingList[3] = Product{name: "Bread", price: 4}

	printShoppingListInfo(shoppingList)
}
