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

			expected := "Contact \"Test_User\" ajouté avec le numéro \"123456\""
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

func setupAnnuaire() {
	annuaire = map[name]phoneNumber{
		"John_Doe": "012345",
		"Jean_luc": "541235",
	}
}
