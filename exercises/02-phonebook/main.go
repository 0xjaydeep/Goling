// Exercise 02 — Phonebook
//
// Problem:
//   Build an in-memory phonebook that stores contacts by name.
//
// Types/functions to implement:
//   type Contact struct { Name, Phone, Email string }
//   addContact(book map[string]Contact, c Contact)
//   findContact(book map[string]Contact, name string) (Contact, bool)
//   deleteContact(book map[string]Contact, name string)
//   listContacts(book map[string]Contact)
//
// Run your solution:  go run main.go
// See the answer:     go run solution.go
package main

import "fmt"

// Contact holds information about a person in the phonebook.
type Contact struct {
	Name  string
	Phone string
	Email string
}

// addContact adds c to book, keyed by c.Name.
func addContact(book map[string]Contact, c Contact) {
	// TODO: store c in book using c.Name as the key
}

// findContact looks up name in book.
// Returns the Contact and true if found, zero Contact and false otherwise.
func findContact(book map[string]Contact, name string) (Contact, bool) {
	// TODO: use the two-value map lookup  →  val, ok := book[name]
	return Contact{}, false
}

// deleteContact removes the contact with the given name from book.
// It is safe to call even if name does not exist.
func deleteContact(book map[string]Contact, name string) {
	// TODO: use the built-in delete function
}

// listContacts prints all contacts in book, one per line.
// If book is empty, print "Phonebook is empty."
func listContacts(book map[string]Contact) {
	// TODO: check if len(book) == 0 and print the empty message
	// TODO: range over book and print each contact's Name, Phone, Email
}

func main() {
	book := make(map[string]Contact)

	// TODO: add these 3 contacts using addContact
	//   Alice  555-0101  alice@example.com
	//   Bob    555-0102  bob@example.com
	//   Carol  555-0103  carol@example.com

	fmt.Println("=== All Contacts ===")
	listContacts(book)

	// TODO: find "Alice" and print her details
	fmt.Println("\n=== Find Alice ===")

	// TODO: delete "Bob", then list again
	fmt.Println("\n=== After Deleting Bob ===")
	listContacts(book)

	// TODO: try to find "Bob" and print a "not found" message
	fmt.Println("\n=== Find Bob (after delete) ===")
}
