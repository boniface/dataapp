package main

import "fmt"

// Singly LL head--> next1--->next2-->tail.next->
//--------------->


// Doubly LL a<==>b<==>c<==>tail.next--> <-prev.newNode.next-> nil
// ------->
// <-------



// Circular  S ( a to b to c to d to a repeat LL
//  --------->
// |       	|
// <--------
// Circular  D ( a to b to c to d to a to d to c repeat LL
//  <--------->
// |       	|
//  <-------->

func main() {
	linkedList()
}

type Switch struct {
	ipAddress string
	ports     int
	name      string
	prev      *Switch
	next      *Switch
}

func linkedList() {
	var head *Switch = new(Switch)
	head.ports = 10
	head.name = "Head Switch"
	head.ipAddress = "1.1.1.1"
	head.next = nil
	head.prev = nil

	var next1 *Switch = &Switch{"1.1.1.2", 20, "Next1", head, nil}
	head.next = next1

	var next2 *Switch = &Switch{"1.1.1.3", 30, "Next2", next1, nil}
	next1.next = next2

	var next3 *Switch = &Switch{"1.1.1.4", 40, "Tail", next2, nil}
	next2.next = next3
	PrintLL(head)
	// Create a New Switch
	var nextSwtch Switch = Switch{"1.1.1.3", 30, "New Switch", nil, nil}
	Add(head,nextSwtch)
	PrintLL(head)
	//Pass it to the Function

	// Print Your New LL

}

func Add(tail *Switch, node Switch)  {
	var newNode *Switch = &Switch{
		node.ipAddress, node.ports, node.name,node.prev, node.next}
	tail.next= newNode
	newNode.prev=tail
}

func InsertNode()  {
	var head *Switch = new(Switch)
	var tail *Switch = new(Switch)
	head.ports = 10
	head.name = "Head Switch"
	head.ipAddress = "1.1.1.1"
	head.prev = nil
	head.next = nil


	var next1 *Switch = &Switch{"1.1.1.2", 20, "Next1", head, nil}
	head.next = next1

	var next2 *Switch = &Switch{"1.1.1.3", 30, "Next2", next1, nil}
	next1.next = next2

	tail.ports = 10
	tail.name = "Tail Node"
	tail.ipAddress = "1.1.1.1"
	tail.prev = next2
	tail.next = nil


}

func insertFunction(position int, node Switch)  {


}

func deleteFunction(){

}


func PrintLL(node *Switch) {

	var n = node
	var end *Switch = nil

	for {
		if n == nil {
			break
		}
		fmt.Printf("%s->", n.name)
		end = n
		n = n.next
	}
	fmt.Printf("End\n\n")

	n = end
	for {
		if n == nil {
			break
		}
		fmt.Printf("%s->", n.name)
		n = n.prev
	}
	fmt.Printf("Head\n\n")
}

// LIFO LAST IN FIRST OUT --> STACK
// FIFO FIRST FIRST OUT --> QUEUE

// ADD, INSERT,  DELETE

// Exercises Practice
// Create Structure called Person { name, prev, next}
// add 5 People
// Link them
// Print going forward and then back ward

// Repeat  for a PC { ipaddre. prev. next } repeat functions as above

// Repeat for a Building{ address , prev, next} repeat the process
