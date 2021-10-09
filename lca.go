package main

import "fmt"

type Node struct {
	data        string
	left, right *Node
}

func main() {

	/**
	*		 	 john
	*	    	/    \
	*	 	pat      mary
	*   	/  \      /   \
	*	 mat sarah  mike  duke
	*
	 */

	x := Node{"john", nil, nil}
	x.left = &Node{"pat", nil, nil}
	x.right = &Node{"mary", nil, nil}

	x.left.left = &Node{"mat", nil, nil}
	x.left.right = &Node{"sarah", nil, nil}

	x.right.left = &Node{"mike", nil, nil}
	x.right.right = &Node{"duke", nil, nil}

	lcaResp := lca(&x, "pat", "mike")
	fmt.Print(lcaResp.data)

}

func lca(head *Node, a string, b string) *Node {

	aExists := pointExistsOnTree(head, a)
	bExists := pointExistsOnTree(head, b)

	/** Ensure points exist inside tree and head exists with data */
	if !aExists || !bExists || len(head.data) == 0 {
		return nil
	}

	/** Search tree for Lowest Common Ancestor */
	return lcaSearch(head, a, b)
}

func lcaSearch(head *Node, a string, b string) *Node {
	if head == nil {
		return nil
	}

	if head.data == a || head.data == b {
		return head
	}

	leftSearch := lcaSearch(head.left, a, b)
	rightSearch := lcaSearch(head.right, a, b)

	if leftSearch != nil {
		if rightSearch != nil {
			return head
		}

		return leftSearch
	}

	return rightSearch
}

func pointExistsOnTree(currentNode *Node, data string) bool {

	/** Ensure root and data intialised */
	if currentNode == nil || len(data) == 0 {
		return false
	}

	/** Check if data is the currentNode */
	if currentNode.data == data {
		return true
	}

	/** Loop recursivly through left and right nodes to check if data exists */
	return pointExistsOnTree(currentNode.left, data) || pointExistsOnTree(currentNode.right, data)

}
