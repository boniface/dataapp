package main

import "fmt"

type Node struct {
	ipAddress string
	ports     int
	name      string
	next      *Node
}

var firstNode *Node = new(Node)
var lastNode  *Node = new(Node)

func main() {
	fmt.Println("Welcome to  ADDING AT POSITION NODE to SINGLY LinkedList")
	createManualList()
	printList(firstNode)
	deleteNode(2)
	printList(firstNode)

}

func createManualList(){
	firstNode.ports = 10
	firstNode.name = "First Node"
	firstNode.ipAddress = "1.1.1.1"
	firstNode.next = nil

	var secondNode *Node = &Node{"1.1.1.2", 20, "Second Node",  nil}

	firstNode.next = secondNode

	var thirdNode *Node = &Node{"1.1.1.3", 30, "Third Node", nil}

	secondNode.next = thirdNode

	var fourthNode *Node = &Node{"1.1.1.4", 40, "Fourth Node", nil}

	thirdNode.next = fourthNode

	lastNode.ports = 10
	lastNode.name = "Fifth Node"
	lastNode.ipAddress = "1.1.1.5"
	lastNode.next = nil

	fourthNode.next=lastNode

}

func deleteNode(position int) {
	var p = firstNode
	var i = 0
	// Move the node to the previous node position that was deleted
	for {
		if p.next == nil || i >= position-1 {
			break
		}
		p = p.next
		i++
	}
	var temp = p.next
	p.next = p.next.next
	temp.next = nil
}


func printList(node *Node) {
	var p = node
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.name)
		p = p.next
	}
	fmt.Printf("End\n\n")
}