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
}
