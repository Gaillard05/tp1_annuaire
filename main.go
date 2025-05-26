package main

import (
	"flag"
	"fmt"
	"os"
	"tp1_annuaire/annuaire"
)

func main() {
	action := flag.String("action", "", "Action : ajouter, rechercher, lister, supprimer, modifier")
	nom := flag.String("nom", "", "Nom du contact")
	prenom := flag.String("prenom", "", "Prénom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone")
	flag.Parse()

	// Charger les données
	if err := annuaire.Charger(); err != nil {
		fmt.Println("Erreur chargement :", err)
		os.Exit(1)
	}

	switch *action {
	case "ajouter":
		if *nom == "" || *prenom == "" || *tel == "" {
			fmt.Println("Erreur : nom, prénom et téléphone requis.")
			return
		}
		err := annuaire.AjouterContact(*nom, *prenom, *tel)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Contact ajouté.")
		}

	case "rechercher":
		if *nom == "" {
			fmt.Println("Erreur : nom requis.")
			return
		}
		c, err := annuaire.RechercherContact(*nom)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Printf("Nom : %s\nPrénom : %s\nTéléphone : %s\n", c.Nom, c.Prenom, c.Tel)
		}

	case "lister":
		liste := annuaire.ListerContacts()
		if len(liste) == 0 {
			fmt.Println("Aucun contact enregistré.")
		} else {
			fmt.Println("Contacts :")
			for _, c := range liste {
				fmt.Printf("- %s %s : %s\n", c.Prenom, c.Nom, c.Tel)
			}
		}

	case "supprimer":
		if *nom == "" {
			fmt.Println("Erreur : nom requis.")
			return
		}
		err := annuaire.SupprimerContact(*nom)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Contact supprimé.")
		}

	case "modifier":
		if *nom == "" || *tel == "" {
			fmt.Println("Erreur : nom et nouveau téléphone requis.")
			return
		}
		err := annuaire.ModifierContact(*nom, *tel)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Contact modifié.")
		}

	default:
		fmt.Println("Action invalide. Utilisez : ajouter, rechercher, lister, supprimer, modifier")
	}
}
