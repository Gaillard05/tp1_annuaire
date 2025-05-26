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
