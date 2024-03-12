package trie

type TRIE struct {
	children map[string]*node
}

type node struct {
	isEnd    bool
	children map[string]*node
}

type Trie interface {
	Lookup(item string) (exists bool)
	stringPresent(item string, checkEnd bool) (exists bool)
	Insert(item string)
	Delete(item string)
	Search(item string) (result []string)
	Display() (result []string)
	IsEmpty() (isEmpty bool)
	isEnd(item string) (isEnd bool)
}

func NewTRIE() Trie {
	return &TRIE{children: map[string]*node{}}
}

func newNode(isEnd bool) *node {
	return &node{isEnd: isEnd, children: map[string]*node{}}
}

func (t *TRIE) Lookup(item string) bool {
	return t.stringPresent(item, true)
}

func (t *TRIE) stringPresent(item string, checkEnd bool) bool {
	if item == "" || t.IsEmpty() {
		return false
	}

	current := t.children[string(item[0])]

	if current == nil {
		return false
	}

	for i := 1; i < len(item); i++ {
		key := string(item[i])
		current = current.children[key]
		if current == nil {
			return false
		}
	}

	if checkEnd {
		return current.isEnd
	}

	return true
}

func (t *TRIE) Insert(item string) {
	if item == "" {
		return
	}

	if t.children[string(item[0])] == nil {
		if len(item) == 1 {
			t.children[string(item[0])] = newNode(true)
			return
		} else {
			t.children[string(item[0])] = newNode(false)
		}
	}

	current := t.children[string(item[0])]

	for i := 1; i < len(item); i++ {
		key := string(item[i])
		if current.children[key] == nil {
			current.children[key] = newNode(false)
		}
		current = current.children[key]
	}
	current.isEnd = true
}

func (t *TRIE) Delete(item string) {
	if item == "" || !t.stringPresent(item, true) {
		return
	}

	current := t.children[string(item[0])]
	backtrack := current
	deleteKey := string(item[0])

	for i := 1; i < len(item); i++ {
		key := string(item[i])
		current = current.children[key]

		if len(current.children) > 1 {
			backtrack = current
			deleteKey = string(item[i+1])
			continue
		}

		if current.isEnd && key != string(item[len(item)-1]) {
			backtrack = current
			deleteKey = string(item[i+1])
		}
	}

	if len(current.children) > 0 {
		current.isEnd = false
		return
	}

	if backtrack == t.children[string(item[0])] {
		children := map[string]*node{}
		for key, node := range t.children {
			if key != deleteKey {
				children[key] = node
			}
		}
		t.children = children
		return
	}

	children := map[string]*node{}
	for key, node := range backtrack.children {
		if key != deleteKey {
			children[key] = node
		}
	}
	backtrack.children = children
}

func (t *TRIE) Search(item string) []string {
	if item == "" || !t.stringPresent(item, false) {
		return []string{}
	}

	current := t.children[string(item[0])]
	for i := 1; i < len(item); i++ {
		key := string(item[i])
		current = current.children[key]
	}

	if len(current.children) == 0 {
		return []string{item}
	}

	var result []string
	searchHelper(current, item, &result)
	return result
}

func (t *TRIE) Display() []string {
	result := []string{}

	if len(t.children) == 0 {
		return result
	}

	for item := range t.children {
		result = append(result, t.Search(item)...)
	}

	return result
}

func (t *TRIE) IsEmpty() bool {
	return len(t.children) == 0
}

func (t *TRIE) isEnd(item string) bool {
	current := t.children[string(item[0])]
	for i := 1; i < len(item); i++ {
		key := string(item[i])
		current = current.children[key]
	}

	return current.isEnd
}

func searchHelper(node *node, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}

	for key, node := range node.children {
		searchHelper(node, prefix+key, result)
	}
}
