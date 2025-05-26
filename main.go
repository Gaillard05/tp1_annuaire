package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp1_annuaire/annuaire"
)

func main() {
	// nom := "anne"
	// c, err := annuaire.RechercherContact(nom)
	// if err != nil {
	// 	fmt.Println("Erreur:", err)
	// 	return
	// }

	// fmt.Println("Nom", c.Nom)
	// fmt.Println("Téléphone", c.Tel)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Rechercher un contact")
		fmt.Println("3. Afficher tous les contacts")
		fmt.Println("4. Quitter")
		fmt.Print("Choix : ")

		scanner.Scan()
		choix := strings.TrimSpace(scanner.Text())

		switch choix {
		case "1":
			fmt.Print("Nom : ")
			scanner.Scan()
			nom := strings.TrimSpace(scanner.Text())

			fmt.Print("Téléphone : ")
			scanner.Scan()
			tel := strings.TrimSpace(scanner.Text())

			_, err := annuaire.CreerContact(nom, tel)

			if err != nil {
				fmt.Println("x", err)
			} else {
				fmt.Println(("Contact ajouté."))
			}
		case "2":
			fmt.Print("Nom à rechercher : ")
			scanner.Scan()
			nom := strings.TrimSpace(scanner.Text())

			c, err := annuaire.RechercherContact(nom)

			if err != nil {
				fmt.Println("x", err)
			} else {
				fmt.Printf("%s - %s\n", c.Nom, c.Tel)
			}
		case "3":
			liste := annuaire.ListerContacts()
			if len(liste) == 0 {
				fmt.Println("Aucun contact enregistré.")
			} else {
				fmt.Println("Liste des contacts :")
				for _, c := range liste {
					fmt.Printf("- %s : %s\n", c.Nom, c.Tel)
				}
			}
		case "4":
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide.")

		}
	}

}
