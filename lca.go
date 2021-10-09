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
