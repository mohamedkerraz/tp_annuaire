# tp_annuaire

**Annuaire CLI - Gestion de contacts en ligne de commande en Go**

Ce programme permet d'ajouter, supprimer, modifier, rechercher et lister des contacts simples 
au format : `Nom → Numéro de téléphone`.

Chaque contact est stocké dans une map Go (`type: map[name]phoneNumber`).

**Commandes disponibles :**

  `--a add`   : Ajoute un nouveau contact avec -n (nom) et -p (téléphone)
  
  `--a del`   : Supprime un contact existant avec -n
  
  `--a put`   : Modifie le numéro d’un contact existant avec -n et -p
  
  `--a get`   : Recherche un contact avec -n
  
  `--a list`  : Affiche tous les contacts existants
  

**Utilisation :**
```
Test add
  go run main.go --a "add" --n "Charlie" --p "0811223344"
Test list
  go run main.go --a "list"
Test get
  go run main.go --a "get" --n "John_Doe"
Test put
  go run main.go --a "put" --n "John_Doe" --p "0600000000"
Test delete
  go run main.go --a "del" --n "John_Doe"
```
**Code :**
```go
// menu choisit l’action à effectuer en fonction des arguments
// et retourne une chaîne affichable, ou termine le programme si erreur
func menu(a action, n name, p phoneNumber) string { ... }

// vérifie si un contact existe dans l’annuaire
func isExists(n name) bool { ... }

// ajoute un nouveau contact si le nom n’existe pas déjà
func addContact(n name, p phoneNumber) string { ... }

// supprime un contact par son nom, s’il existe
func deleteContact(n name) string { ... }

// modifie le numéro d’un contact existant
func putContact(n name, p phoneNumber) string { ... }

// retourne le numéro d’un contact s’il existe
func getContact(n name) string { ... }

// retourne tous les contacts sous forme de chaîne formatée
func listContact() string { ... }
```
