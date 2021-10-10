package lca

import (
	"testing"
)

func TestLCA(t *testing.T) {

	/**
	*		 	 john
	*	    	/    \
	*	 	pat      mary
	*
	 */

	x := Node{"john", nil, nil}
	x.left = &Node{"pat", nil, nil}
	x.right = &Node{"mary", nil, nil}

	lcaResp := lca(&x, "pat", "mary")

	if lcaResp.data != "john" {
		t.Errorf(`lca(Node, "pat", "mary") = %s; wants "john"`, lcaResp.data)
	}
}

func TestMissingReference(t *testing.T) {
	/**
	*		 	 john
	*	    	/    \
	*	 	pat      mary
	*
	 */

	x := Node{"john", nil, nil}
	x.left = &Node{"pat", nil, nil}
	x.right = &Node{"mary", nil, nil}

	lcaResp := lca(&x, "pat", "doesNotExistOnTree")

	if lcaResp != nil {
		t.Errorf(`lca(Node, "pat", "mike") = %s; wants nil`, lcaResp.data)
	}
}

