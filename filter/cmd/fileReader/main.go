package main

import (
	"filter/int/pkg"
	"fmt"
	"log"
)

func main() {

	file, err := pkg.LoadFile()
	if err != nil {
		log.Fatalf("Failed to load file %s", err)
	}

	txtLines, err := pkg.Scanner(file)
	if err != nil {
		log.Fatalf("Failed to scan file %s", err)
	}

	for _, txtLine := range txtLines {
		fmt.Println(txtLine)
	}

}
