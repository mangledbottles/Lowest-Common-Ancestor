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
}