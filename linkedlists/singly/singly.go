package main

import "fmt"

type Node struct {
	ipAddress string
	ports     int
	name      string
	next      *Node
}
var firstNode *Node = new(Node)


func main() {
	fmt.Println("Welcome to SINGLY Linked List")
	createManualList()
	printList(firstNode)
	printNodes(firstNode)
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

func printNodes(node *Node) {
	var p = node
	for {
		if p == nil {
			fmt.Println(" Why Iam Breaking ", p)
			break
		}
		fmt.Println("The Node ", p.name, " has ", p.ipAddress, " and Port ", p.ports)
		p = p.next
	}

}