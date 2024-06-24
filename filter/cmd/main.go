package main

import (
	"bufio"
	"filter/int/pkg"
	"filter/int/util"
	"fmt"
	"os"
)

func main() {
	filepath := ""

	content, err := util.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error opening file:- %s", filepath)
	}

	fmt.Print("Enter the search string....")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if util.ValidateString(input) {
		count := pkg.Filter(content, input)
		fmt.Printf("Search string %s found %d times .\n", input, count)
	} else {
		fmt.Printf("Invalid input string")
	}
}
