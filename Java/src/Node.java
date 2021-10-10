package Java.src;

public class Node {
    public final String data;
    public Node left, right;

    public Node(String data, Node left, Node right) {
        this.data = data;
        this.left = left;
        this.right = right;
    }
}