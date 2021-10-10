package main
import "fmt"
type Node struct {
   value int
   next *Node
}
func NewNode(value int) *Node{
   var n Node
   n.value = value
   n.next = nil
   return &n
}
func TraverseLinkedList(head *Node){
   fmt.Printf("Linked List: ")
   temp := head
   for temp != nil {
      fmt.Printf("%d ", temp.value)
      temp = temp.next
   }
}
func main(){
   head := NewNode(30)
   head.next = NewNode(10)
   head.next.next = NewNode(40)
   head.next.next.next = NewNode(40)
   TraverseLinkedList(head)
}