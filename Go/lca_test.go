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

func TestLargeTree(t *testing.T) {
	/**
	* Initialise tree structure with data
	*
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

	/** Initialise answer table for testing */
	lcaAnswerTable := []struct {
		predecessor1   string
		predecessor2   string
		commonAncestor string
	}{
		{"pat", "mary", "john"},
		{"mat", "sarah", "pat"},
		{"mike", "duke", "mary"},
		{"mat", "pat", "pat"},
		{"mary", "sarah", "john"},
	}

	/** Test LCA algorithm against answer table */
	for _, lcaAnswer := range lcaAnswerTable {
		lcaResp := lca(&x, lcaAnswer.predecessor1, lcaAnswer.predecessor2)

		if lcaResp.data != lcaAnswer.commonAncestor {
			t.Errorf(`lca(Node, "%s", "%s") = %s; wants "%s"`,
				lcaAnswer.predecessor1,
				lcaAnswer.predecessor2,
				lcaResp.data,
				lcaAnswer.commonAncestor,
			)

		}
	}
}

func TestSingularTree(t *testing.T) {
	/**
	*
	*	tom
	*
	 */

	x := Node{"tom", nil, nil}

	lcaResp := lca(&x, "tom", "tom")

	if lcaResp.data != "tom" {
		t.Errorf(`lca(Node, "tom", "tom") = %s; wants "tom"`, lcaResp.data)
	}

}

func TestSimpleDAGTree(t *testing.T) {
	/**
	*
	*     a
	*    / \
	*   b   c
	*    \ /
	*	  d
	*	  |
	*	  e
	 */

	x := Node{"a", nil, nil}
	commonNode := &Node{"d", nil, nil}

	x.left = &Node{"b", nil, nil}
	x.left.right = commonNode

	x.right = &Node{"c", nil, nil}
	x.right.left = commonNode

	commonNode.left = &Node{"e", nil, nil}
}

func TestDAGTree(t *testing.T) {
	/**
	*
	*     a
	*    / \
	*   b   c
	*    \ /
	*	  d
	 */

	x := Node{"a", nil, nil}
	x.left = &Node{"b", nil, nil}
	x.left.right = &Node{"d", nil, nil}

	x.right = &Node{"c", nil, nil}
	x.right.left = x.left.right
}
