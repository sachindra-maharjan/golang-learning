package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Sachindra", "Maharjan")
	err := p.SetTwitterHandler("@sachindramaharjan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("ID: %v Name: %v Twitter Handle: %v\n", p.ID(), p.FullName(), p.TwitterHandle())
}