package sensitive_words_filter


// DFA node
type Node struct {
	End bool
	Next map[rune]*Node
}

// Add node
func (n *Node) AddChild(c rune) *Node {
	if n.Next == nil {
		n.Next = make(map[rune]*Node)
	}

	if next, ok := n.Next[c]; ok {
		return next
	} else {
		n.Next[c] = &Node{
			End:  false,
			Next: nil,
		}
		return n.Next[c]
	}
}

// Find node
func (n *Node) FindChild(c rune) *Node {
	if n.Next == nil {
		return nil
	}

	if _, ok := n.Next[c]; ok {
		return n.Next[c]
	}
	return nil
}

func (n *Node) AddWord(word string) {
	node := n
	chars := []rune(word)
	for index, _ := range chars {
		node = node.AddChild(chars[index])
	}
	node.End = true
}