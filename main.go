package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addContact(phoneBook map[string]string, scanner *bufio.Scanner) {
	fmt.Println("Enter the name ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter the phone number ")
	scanner.Scan()
	PhoneNmber := scanner.Text()
	phoneBook[name] = PhoneNmber

	fmt.Println("Contact Added Successfully ")

}
func viewContacts(phoneBook map[string]string) {
	fmt.Println("The Contacts Include ")
	fmt.Println(phoneBook)

}
func searchContact(phoneBook map[string]string, scanner *bufio.Scanner) {
	fmt.Println("Enter the name you want to search ")
	scanner.Scan()
	search := strings.TrimSpace(scanner.Text())
	if number, exists := phoneBook[search]; exists {
		fmt.Printf("Found! %s: %s\n", search, number)
	} else {
		fmt.Println("The name is not present in the directory.")
	}
}
func deleteContact(phoneBook map[string]string, scanner *bufio.Scanner) {
	fmt.Println("Enter the name you want to delete ")
	scanner.Scan()
	deleted := strings.TrimSpace(scanner.Text())
	if _, exists := phoneBook[deleted]; exists {
		delete(phoneBook, deleted)
		fmt.Println("Deleted Successfully ")
	} else {
		fmt.Println("Contact not found.")
	}

}

func main() {
	phoneBook := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Phone Book CLI")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contact")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Exit")
		fmt.Println("5. Delete Contact")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			addContact(phoneBook, scanner)
		case "2":
			viewContacts(phoneBook)
		case "3":
			searchContact(phoneBook, scanner)
		case "4":
			fmt.Println("Exiting Phone Book CLI. Goodbye!")
			return
		case "5":
			deleteContact(phoneBook, scanner)
		default:
			fmt.Println("Invalid choice! Please enter a valid option.")
		}
	}
}
