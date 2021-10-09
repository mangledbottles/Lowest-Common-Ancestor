package lca

import (
	"fmt"
	"testing"
)

func TestLCA(t *testing.T) {

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

	if lcaResp.data != "john" {
		t.Errorf(`lca(Node, "pat", "mike") = %s; wants "john"`, lcaResp.data)
	}
	fmt.Print(lcaResp.data)

}
