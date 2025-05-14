package main

import (
	"flag"
	"fmt"
)

func main() {
	// Déclaration des variables
	var a action
	var n name
	var p phone_number

	flag.StringVar((*string)(&a), "a", "", "choose action")

	flag.StringVar((*string)(&n), "n", "", "choose a contact name")

	flag.StringVar((*string)(&p), "p", "", "put a number")
	fmt.Println(a, n, p)
	flag.Parse()

	res := menu(a, n, p)

	fmt.Println(res)
}

func menu(a action, n name, p phone_number) string {
	switch a {
	case "add":
		return add_contact(n, p)
	case "del":
		return delete_contact(n)
	// case "3":
	// 	put_contact(n)
	case "get":
		return get_contact(n)
	case "list":
		return list_contact()

	}
	return ""
}

type name string
type action string

type phone_number string

var annuaire = map[name]phone_number{
	"John_Doe": "012345",
	"Jean_luc": "541235",
}

func isExists(n name) bool {

	_, value := annuaire[n]
	return value

}

func add_contact(n name, p phone_number) string {

	if isExists(n) {
		return "contact already exists"
	}
	annuaire[n] = p
	return fmt.Sprintf("Contact %q ajouté avec le numéro %q", n, p)

}

func list_contact() string {
	var result string
	result = "Liste des contacts :\n"
	for numero, nom := range annuaire {
		result += fmt.Sprintf("- %s : %s\n", numero, nom)
	}
	return result
}

func get_contact(n name) string {

	number, _ := annuaire[n]

	if isExists(n) {
		return fmt.Sprintf("contact number: %q ", number)
	}
	return "contact not found"
}

func delete_contact(n name) string {
	if isExists(n) {
		delete(annuaire, n)
		return fmt.Sprintf("contact : %q has been deleted", n)
	}
	return "contact not found"
}
