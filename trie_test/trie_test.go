package trie_test

import (
	"bhargav/trie/trie"
	"testing"
)

func TestTRIELookup(test *testing.T) {
	t := trie.NewTRIE()
	t.Insert("test")

	if !t.Lookup("test") {
		test.Error("Expected to find 'test' in TRIE")
	}

	if t.Lookup("testing") {
		test.Error("Expected not to find 'testing' in TRIE")
	}

	if t.Lookup("tes") {
		test.Error("Expected not to find 'tes' in TRIE")
	}

	if t.Lookup("te") {
		test.Error("Expected not to find 'te' in TRIE")
	}

	if t.Lookup("t") {
		test.Error("Expected not to find 'te' in TRIE")
	}

	if t.Lookup("") {
		test.Error("Expected not to find empty string in TRIE")
	}
}

func TestTRIEInsert(test *testing.T) {
	t := trie.NewTRIE()
	t.Insert("test")

	if !t.Lookup("test") {
		test.Error("Expected to find 'test' in TRIE after insertion")
	}

	if t.IsEmpty() {
		test.Error("Expected TRIE not to be empty after insertion")
	}
}

func TestTRIEDelete(test *testing.T) {
	t := trie.NewTRIE()
	t.Insert("test")
	t.Delete("test")

	if t.Lookup("test") {
		test.Error("Expected not to find 'test' in TRIE after deletion")
	}

	if !t.IsEmpty() {
		test.Error("Expected TRIE to be empty after deletion")
	}
}

func TestTRIESearch(test *testing.T) {
	t := trie.NewTRIE()
	t.Insert("test")
	t.Insert("testing")
	t.Insert("testcase")

	expected := []string{"test", "testing", "testcase"}
	result := t.Search("test")

	for _, v := range expected {
		if !listContainsItem(v, result) {
			test.Errorf("Expected %s not found in result: %v", v, result)
		}
	}
}

func TestTRIEDisplay(test *testing.T) {
	t := trie.NewTRIE()

	result := t.Display()
	if len(result) != 0 {
		test.Error("Expected TRIE to be empty")
	}

	t.Insert("test")
	t.Insert("testing")
	t.Insert("testcase")

	expected := []string{"test", "testing", "testcase"}
	result = t.Display()
	for _, v := range expected {
		if !listContainsItem(v, result) {
			test.Errorf("Expected %s not found in result: %v", v, result)
		}
	}
}

func TestTRIEIsEmpty(test *testing.T) {
	t := trie.NewTRIE()

	if !t.IsEmpty() {
		test.Error("Expected TRIE to be empty")
	}

	t.Insert("test")

	if t.IsEmpty() {
		test.Error("Expected TRIE not to be empty after insertion")
	}
}

func listContainsItem(item string, list []string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
