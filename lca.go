package main

import "fmt"

type Node struct {
	data        string
	left, right *Node
}

func main() {
	/**
	*		 john
	*	    /    \
	*	 pat      mary
	*   /  \      /   \
	* mat sarah  mike  duke
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

}

func traverseTree(n Node, index int) {
	index = index + 1
	var toPrint string = ""
	if n.left != nil {
		// toPrint
		traverseTree(*n.left, index)
	}

	if n.right != nil {

		traverseTree(*n.right, index)
	}

	fmt.Println(toPrint)

}
