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
        System.out.println(head);

        return new Node("john", null, null);
    }

    private Node LcaRecursive(Node currentNode) {
        return new Node("john", null, null);
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