package main

import (
	"fmt"

	"github.com/johnwesonga/go-mzalendo/mzalendo"
)

func main() {
	client := mzalendo.NewClient(nil)
	results, err := client.Persons.GetPerson("1290")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Print(results)

}
