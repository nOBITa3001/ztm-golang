//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func showLine(lines []Part) {
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}

func main() {
	assemblyLines := make([]Part, 3)
	assemblyLines[0] = "Pipe"
	assemblyLines[1] = "Bolt"
	assemblyLines[2] = "Sheet"

	fmt.Println("3 parts:")
	showLine(assemblyLines)

	assemblyLines = append(assemblyLines, "Washer", "Wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyLines)

	assemblyLines = assemblyLines[3:]
	fmt.Println("\nSliced:")
	showLine(assemblyLines)
}
