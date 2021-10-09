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

	fmt.Print(x)
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
