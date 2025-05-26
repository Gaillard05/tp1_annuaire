package main

import (
	"fmt"
	"tp1_annuaire/annuaire"
)

func main() {
	nom := "anne"
	c, err := annuaire.RechercherContact(nom)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	fmt.Println("Nom", c.Nom)
	fmt.Println("Téléphone", c.Tel)
}
