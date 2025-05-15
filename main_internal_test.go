package main

import (
	"fmt"
	"testing"
)

func TestAddContact(t *testing.T) {
	t.Run("New contact scenarios", func(t *testing.T) {
		setupAnnuaire()

		t.Run("should add a new contact successfully", func(t *testing.T) {
			testName := name("Test_User")
			testPhone := phoneNumber("123456")
			result := addContact(testName, testPhone)

			expected := "Contact \"Test_User\" added with number \"123456\""
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}

			if phone, exists := annuaire[testName]; !exists || phone != testPhone {
				t.Errorf("Contact not properly added to annuaire")
			}
		})

		t.Run("should not add a contact that already exists", func(t *testing.T) {
			existingName := name("John_Doe")
			originalPhone := annuaire[existingName]
			newPhone := phoneNumber("999999")

			result := addContact(existingName, newPhone)

			expected := "contact already exists"
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}

			if annuaire[existingName] != originalPhone {
				t.Errorf("Existing contact was modified")
			}
		})
	})

}

func TestGetContact(t *testing.T) {
	t.Run("Get contact number scenarios", func(t *testing.T) {
		setupAnnuaire()

		t.Run("should retrieve an existing contact successfully", func(t *testing.T) {
			existingName := name("John_Doe")
			expectedNumber := annuaire[existingName]

			result := getContact(existingName)

			expected := fmt.Sprintf("contact number: %q ", expectedNumber)
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}
		})

		t.Run("should return not found for non-existent contact", func(t *testing.T) {
			nonExistingName := name("Nobody_Here")

			result := getContact(nonExistingName)

			expected := "contact not found"
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}
		})

		t.Run("should handle empty name parameter", func(t *testing.T) {
			emptyName := name("")

			result := getContact(emptyName)

			expected := "contact not found"
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}
		})
	})
}

func TestPutContact(t *testing.T) {
	t.Run("Put contact scenarios", func(t *testing.T) {
		setupAnnuaire()

		t.Run("should update an existing contact's number", func(t *testing.T) {
			existingName := name("John_Doe")
			newPhone := phoneNumber("777888")

			result := putContact(existingName, newPhone)

			expected := fmt.Sprintf("the number of %q was updated: %q", existingName, newPhone)
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}

			if annuaire[existingName] != newPhone {
				t.Errorf("Contact number was not updated correctly in annuaire")
			}
		})

		t.Run("should fail to update a non-existent contact", func(t *testing.T) {
			nonExistingName := name("Ghost_User")
			newPhone := phoneNumber("000000")

			result := putContact(nonExistingName, newPhone)

			expected := fmt.Sprintf("contact not found", nonExistingName)
			if result != expected {
				t.Errorf("Expected: %q, got: %q", expected, result)
			}

			if _, exists := annuaire[nonExistingName]; exists {
				t.Errorf("Non-existent contact was wrongly added to annuaire")
			}
		})
	})
}


func setupAnnuaire() {
	annuaire = map[name]phoneNumber{
		"John_Doe": "012345",
		"Jean_luc": "541235",
	}
}
