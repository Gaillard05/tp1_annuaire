package annuaire

import (
	"encoding/json"
	"errors"
	"os"
)

type Contact struct {
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Tel    string `json:"tel"`
}

var contacts = make(map[string]Contact)

const fichier = "contacts.json"

// Chargement
func Charger() error {
	data, err := os.ReadFile(fichier)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &contacts)
}

// Sauvegarde
func sauvegarder() error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fichier, data, 0644)
}

// Ajouter
func AjouterContact(nom, prenom, tel string) error {
	if _, existe := contacts[nom]; existe {
		return errors.New("contact déjà existant")
	}
	contacts[nom] = Contact{Nom: nom, Prenom: prenom, Tel: tel}
	return sauvegarder()
}

// Rechercher
func RechercherContact(nom string) (Contact, error) {
	if c, ok := contacts[nom]; ok {
		return c, nil
	}
	return Contact{}, errors.New("contact non trouvé")
}

// Lister
func ListerContacts() []Contact {
	liste := []Contact{}
	for _, c := range contacts {
		liste = append(liste, c)
	}
	return liste
}

// Supprimer
func SupprimerContact(nom string) error {
	if _, ok := contacts[nom]; !ok {
		return errors.New("contact non trouvé")
	}
	delete(contacts, nom)
	return sauvegarder()
}

// Modifier
func ModifierContact(nom, nouveauTel string) error {
	c, ok := contacts[nom]
	if !ok {
		return errors.New("contact non trouvé")
	}
	c.Tel = nouveauTel
	contacts[nom] = c
	return sauvegarder()
}
