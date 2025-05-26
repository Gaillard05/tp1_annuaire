#  TP1 - Annuaire en Go

Ce projet est une application en ligne de commande développée en Go pour gérer un annuaire simple.


## Fonctionnalités

- Ajouter un contact (`nom`, `prenom`, `tel`)
- Rechercher un contact par nom
- Lister tous les contacts
- Supprimer un contact
- Modifier un contact existant
- Vérifier l’unicité du nom
- Sauvegarde automatique dans un fichier "contacts.json"

---

### Compilation / Exécution directe :

Aller dans le répertoire cd tp1_annuaire 

Lancez les commandes suivante :

Ajouter un contact :
go run main.go --action ajouter --nom "Charlie" --prenom "Duval" --tel "0811223344"

Rechercher un contact :
go run main.go --action rechercher --nom "Charlie"

Afficher la liste des contacts :
go run main.go --action lister

Modifier un contact :
go run main.go --action modifier --nom "Charlie" --tel "0612345678"

Supprimer un contact :
go run main.go --action supprimer --nom "Charlie"