//go:build ignore

// Solution for Exercise 02 — Phonebook
// Run with: go run solution.go
package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

func addContact(book map[string]Contact, c Contact) {
	book[c.Name] = c
}

func findContact(book map[string]Contact, name string) (Contact, bool) {
	c, ok := book[name]
	return c, ok
}

func deleteContact(book map[string]Contact, name string) {
	delete(book, name)
}

func listContacts(book map[string]Contact) {
	if len(book) == 0 {
		fmt.Println("  Phonebook is empty.")
		return
	}
	for _, c := range book {
		fmt.Printf("  %-8s  %-12s  %s\n", c.Name, c.Phone, c.Email)
	}
}

func main() {
	book := make(map[string]Contact)

	addContact(book, Contact{"Alice", "555-0101", "alice@example.com"})
	addContact(book, Contact{"Bob", "555-0102", "bob@example.com"})
	addContact(book, Contact{"Carol", "555-0103", "carol@example.com"})

	fmt.Println("=== All Contacts ===")
	listContacts(book)

	fmt.Println("\n=== Find Alice ===")
	if c, ok := findContact(book, "Alice"); ok {
		fmt.Printf("  Found: %+v\n", c)
	}

	fmt.Println("\n=== After Deleting Bob ===")
	deleteContact(book, "Bob")
	listContacts(book)

	fmt.Println("\n=== Find Bob (after delete) ===")
	if _, ok := findContact(book, "Bob"); !ok {
		fmt.Println("  Bob not found in phonebook.")
	}
}
