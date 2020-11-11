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
	var newNode *Node = &Node{"1.1.1.20", 70, "Sixth Node",  nil, nil}
	PrintDoubleCircularList()
	insertAtPosition(3, newNode)
	PrintDoubleCircularList()

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
	lastNode.name = "Last Node"
	lastNode.ipAddress = "1.1.1.5"
	lastNode.prev = fourthNode
	lastNode.next = firstNode

	fourthNode.next = lastNode
	firstNode.prev =lastNode
}


func insertAtPosition(insertPosition int, newNode *Node) {
	var p = firstNode
	var i = 0
	for {
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}
	newNode.next = p.next // newNode next point to next node
	p.next = newNode // current next point to newNode
	newNode.prev = p
	newNode.next.prev = newNode
}

func PrintDoubleCircularList() {
	var p = firstNode
	for {
		fmt.Printf("%s -> ", p.name)
		p = p.next
		if p == firstNode {
			break
		}
	}
	fmt.Printf("%s ", p.name)
	fmt.Printf("End\n")
	p = lastNode
	for {
		fmt.Printf("%s -> ", p.name)
		p = p.prev
		if p == lastNode {
			break
		}
	}
	fmt.Printf("%s ", p.name)
	fmt.Printf("Start\n\n")
}



