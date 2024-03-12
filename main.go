package main

import (
	"bhargav/trie/trie"
	"bufio"
	"fmt"
	"os"
)

func main() {
	t := trie.NewTRIE()

	for {
		var userInput string

		fmt.Println("---MENU---")
		fmt.Println("1. Insert Item to Trie")
		fmt.Println("2. Lookup for Keyword in Trie")
		fmt.Println("3. Search for Items in Trie by Keyword")
		fmt.Println("4. Delete Item from Trie")
		fmt.Println("5. Display Items in Trie")
		fmt.Println("6. Exit")
		fmt.Print("Choice: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userInput = scanner.Text()
		} else {
			continue
		}
		fmt.Println()

		switch userInput {
		case "1":
			var item string
			fmt.Print("Enter Item for Insertion: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				item = scanner.Text()
			} else {
				continue
			}
			t.Insert(item)
		case "2":
			var item string
			fmt.Print("Enter Keyword for Lookup: ")
			if scanner.Scan() {
				item = scanner.Text()
			} else {
				continue
			}
			if t.Lookup(item) {
				fmt.Println(item, "is present")
			} else {
				fmt.Println(item, "is not present")
			}
		case "3":
			var item string
			fmt.Print("Enter Keyword for Search: ")
			if scanner.Scan() {
				item = scanner.Text()
			} else {
				continue
			}
			displaySearchResults(item, t.Search(item))
		case "4":
			var item string
			fmt.Print("Enter Item for Deletion: ")
			if scanner.Scan() {
				item = scanner.Text()
			} else {
				continue
			}
			t.Delete(item)
		case "5":
			res := t.Display()
			displayResult(res)
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
		fmt.Println()
	}
}

func displaySearchResults(keyword string, result []string) {
	if len(result) == 0 {
		fmt.Println("No item found for", keyword)
	} else {
		fmt.Printf("Items matching %v are:\n", keyword)
		for i, item := range result {
			fmt.Printf("%v: %v\n", i+1, item)
		}
	}
}

func displayResult(result []string) {
	if len(result) == 0 {
		fmt.Println("No items")
		return
	}

	fmt.Println("Items are:")
	for i, item := range result {
		fmt.Printf("%v: %v\n", i+1, item)
	}
}
