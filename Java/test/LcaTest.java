package Java.test;

import Java.src.Lca;
import Java.src.Node;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import static org.junit.Assert.*;

@RunWith(JUnit4.class)
public class LcaTest {
    @Test
    public void TestLCA() {
        /* Initialise tree structure with data
         		 	 john
         	    	/    \
         	 	 pat      mary
         */

        Node n = new Node("john", null, null);
        n.left = new Node("pat", null, null);
        n.right = new Node("mary", null, null);

        Node lcaResp = new Lca(n, "pat", "mary").FindLca();

        assertEquals("lca(Node, pat, mary) = john; wants john", "john", lcaResp.data);

        System.out.println(lcaResp.data);
    }

    @Test
    public void TestLargeTree() {
        /*
         * Initialise tree structure with data
         *
         *		 	 john
         *	    	/    \
         *	 	pat      mary
         *   	/  \      /   \
         *	 mat sarah  mike  duke
         *
         */

        Node n = new Node("john", null, null);
        n.left = new Node("pat", null, null);
        n.right = new Node("mary", null, null);

        n.left.left = new Node("mat", null, null);
        n.left.right = new Node("sarah", null, null);

        n.right.left = new Node("mike", null, null);
        n.right.right = new Node("duke", null, null);

        String[][] lcaAnswerTable = new String[][] {
                {"pat", "mary", "john"},
                {"mat", "sarah", "pat"},
                {"mike", "duke", "mary"},
                {"mat", "pat", "john"},
                {"mary", "sarah", "john"},
        };

        String predecessor1, predecessor2, commonAncestor;
        for (String[] lcaAnswer : lcaAnswerTable) {
            predecessor1 = lcaAnswer[0];
            predecessor2 = lcaAnswer[1];
            commonAncestor = lcaAnswer[2];

            Node lcaResp = new Lca(n, predecessor1, predecessor2).FindLca();
            assertEquals(String.format("lca(Node, %s, %s) = %s; wants %s", predecessor1, predecessor2, lcaResp, commonAncestor),
                    commonAncestor, lcaResp.data);

        }
}
