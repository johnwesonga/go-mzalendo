package main

import (
	"fmt"

	"github.com/johnwesonga/go-mzalendo/mzalendo"
)

func main() {
	client := mzalendo.NewClient(nil)
	r, err := client.Api.GetPerson("1290")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(r.Result.Name)

	org, err := client.Api.GetOrganization("148")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(org.Result.Name)

}
