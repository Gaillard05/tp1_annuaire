package annuaire

import "fmt"

type ListContact string

type Contact struct {
	Nom string
	Tel string
}

var contacts = map[string]Contact{
	"anne": {Nom: "Anne", Tel: "0680201823"},
}

func RechercherContact(nom string) (Contact, error) {
	c, ok := contacts[nom]
	if !ok {
		return Contact{}, fmt.Errorf("aucun contact trouvé pour %q", nom)
	}

	return c, nil
}

func ListerContacts() []Contact {
	var liste []Contact
	for _, c := range contacts {
		liste = append(liste, c)
	}
	return liste
}

func CreerContact(nom, tel string) (Contact, error) {
	if _, existe := contacts[nom]; existe {
		return Contact{}, fmt.Errorf("le contact %q existe déjà", nom)
	}

	c := Contact{Nom: nom, Tel: tel}
	contacts[nom] = c
	return c, nil
}
