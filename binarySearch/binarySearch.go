package binarySearch

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

//BST String()
//In order traversal

// left itself right

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inOrderTraversalByNode(sb, root.right)
}

// func main() {
// 	// n := &node{1, nil, nil}
// 	// n.left = &node{0, nil, nil}
// 	// n.right = &node{2, nil, nil}

// 	// b := bst{
// 	// 	root: n,
// 	// 	len:  3,
// 	// }

// 	// fmt.Println(b)

// }
