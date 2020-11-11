package main

import "fmt"

type Node struct {
	ipAddress string
	ports     int
	name      string
	prev      *Node
	next      *Node
}

var firstNode *Node = new(Node)
var lastNode *Node = new(Node)

func main() {
	createManually()
	printList(firstNode)

}

func createManually(){
	firstNode.ports = 10
	firstNode.name = "First Node"
	firstNode.ipAddress = "1.1.1.1"
	firstNode.prev = nil
	firstNode.next = nil

	var secondNode *Node = &Node{"1.1.1.2", 20, "Second Node",  firstNode, nil}

	firstNode.next = secondNode

	var thirdNode *Node = &Node{"1.1.1.3", 30, "Third Node", secondNode, nil}

	secondNode.next = thirdNode

	var fourthNode *Node = &Node{"1.1.1.4", 40, "Fourth Node", thirdNode, nil}

	thirdNode.next = fourthNode

	lastNode.ports = 50
	lastNode.name = "Fifth Node"
	lastNode.ipAddress = "1.1.1.5"
	lastNode.prev = fourthNode
	lastNode.next = nil

	fourthNode.next = lastNode
}

func printList(node *Node) {
	var p = node
	var end *Node = nil
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.name)
		end = p
		p = p.next
	}
	fmt.Printf("End\n")
	p = end
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.name)
		p = p.prev
	}
	fmt.Printf("Start\n\n")
}







