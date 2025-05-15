package main

import (
	"flag"
	"fmt"
)

func main() {
	// Déclaration des variables
	var a action
	var n name
	var p phoneNumber

	flag.StringVar((*string)(&a), "a", "", "choose action")

	flag.StringVar((*string)(&n), "n", "", "choose a contact name")

	flag.StringVar((*string)(&p), "p", "", "put a number")
	fmt.Println(a, n, p)
	flag.Parse()

	res := menu(a, n, p)

	fmt.Println(res)
}

func menu(a action, n name, p phoneNumber) string {
	switch a {
	case "add":
		return addContact(n, p)
	case "del":
		return deleteContact(n)
	// case "3":
	// 	put_contact(n)
	case "get":
		return getContact(n)
	case "list":
		return listContact()

	}
	return ""
}

type name string
type action string

type phoneNumber string

var annuaire = map[name]phoneNumber{
	"John_Doe": "012345",
	"Jean_luc": "541235",
}

func isExists(n name) bool {

	_, value := annuaire[n]
	return value

}

func addContact(n name, p phoneNumber) string {

	if isExists(n) {
		return "contact already exists"
	}
	annuaire[n] = p
	return fmt.Sprintf("Contact %q ajouté avec le numéro %q", n, p)

}

func listContact() string {
	var result string
	result = "Liste des contacts :\n"
	for numero, nom := range annuaire {
		result += fmt.Sprintf("- %s : %s\n", numero, nom)
	}
	return result
}

func getContact(n name) string {

	number, _ := annuaire[n]

	if isExists(n) {
		return fmt.Sprintf("contact number: %q ", number)
	}
	return "contact not found"
}

func deleteContact(n name) string {
	if isExists(n) {
		delete(annuaire, n)
		return fmt.Sprintf("contact : %q has been deleted", n)
	}
	return "contact not found"
}
