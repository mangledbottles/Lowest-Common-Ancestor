package Java.src;

public class Lca {
    public Node head;
    public String a, b;

    public Lca(Node head, String a, String b) {
        this.head = head;
        this.a = a;
        this.b = b;
    }

    public Node FindLca() {

        boolean aExists = pointExistsOnTree(head, a);
        boolean bExists = pointExistsOnTree(head, b);

        /* Ensure points exist inside tree and head exists with data */
        if(!aExists || !bExists || head.data.length() == 0) return null;

        /* Search tree for Lowest Common Ancestor */
        return LcaRecursive(head);
    }

    private Node LcaRecursive(Node currentNode) {
        if (currentNode == null)
            return null;

        if (currentNode.data.equals(a) || currentNode.data.equals(b))
            return head;

        Node leftSearch = LcaRecursive(currentNode.left);
        Node rightSearch = LcaRecursive(currentNode.right);

        if (leftSearch != null) {
            if (rightSearch != null)
                return currentNode;

            return leftSearch;
        }

        return rightSearch;
    }

    private boolean pointExistsOnTree(Node currentNode, String data) {
        /* Ensure root and data initialised */
        if(currentNode == null || data.length() == 0) return false;
        /* Check if data is the currentNode */
        if(currentNode.data.equals(data)) return true;

        /* Loop recursively through left and right nodes to check if data exists */
        return pointExistsOnTree(currentNode.left, data) || pointExistsOnTree(currentNode.right, data);
    }

}