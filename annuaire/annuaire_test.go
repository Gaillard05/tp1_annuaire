package annuaire

import (
	"os"
	"testing"
)

// Utilitaire : reset l'annuaire avant chaque test
func resetAnnuaire() {
	contacts = make(map[string]Contact)
	os.Remove(fichier) // Supprime le fichier de sauvegarde s’il existe
}

// Ajouter un contact, puis le rechercher
func TestAjouterEtRechercherContact(t *testing.T) {
	resetAnnuaire()

	err := AjouterContact("TestNom", "Alice", "0102030405")
	if err != nil {
		t.Fatalf("Erreur à l'ajout : %v", err)
	}

	c, err := RechercherContact("TestNom")
	if err != nil {
		t.Fatalf("Erreur à la recherche : %v", err)
	}

	if c.Prenom != "Alice" || c.Tel != "0102030405" {
		t.Errorf("Contact incorrect. Attendu: Alice 0102030405, Obtenu: %s %s", c.Prenom, c.Tel)
	}
}

// Modifier un contact existant
func TestModifierContact(t *testing.T) {
	resetAnnuaire()

	AjouterContact("Dupont", "Jean", "0600000000")
	err := ModifierContact("Dupont", "0699999999")
	if err != nil {
		t.Fatalf("Erreur à la modification : %v", err)
	}

	c, err := RechercherContact("Dupont")
	if err != nil {
		t.Fatalf("Erreur à la recherche après modif : %v", err)
	}

	if c.Tel != "0699999999" {
		t.Errorf("Téléphone non modifié. Attendu: 0699999999, Obtenu: %s", c.Tel)
	}
}
