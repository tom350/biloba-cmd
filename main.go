package main

import (
	"biloba-cmd/driver"
	"biloba-cmd/handler"
	"fmt"
	"os"
)

func main() {
	UserID := os.Args[1]

	db, _ := driver.GetDBClient()
	kids := handler.KidPrint(&handler.App{&db}, UserID)

	var s string

	for id, kid := range kids {
		fmt.Print(id+1)
		fmt.Println("- Son ID est :", kid.ID)
		fmt.Println("Son nom est :", kid.Name)
		fmt.Println("Date de naissance :", kid.BirthDate)
		fmt.Println("Ville :", kid.City)
		if kid.Gender == 0 {s = "fille"}
		if kid.Gender == 1 {s = "garçon"}
		fmt.Println("C'est un(e) " + s)
		fmt.Println()
	}

}