package main

import "fmt"

type Switch struct {
	ipAddress string
	ports     int
	name      string
	prev      *Switch
	next      *Switch
}


func main() {
	fmt.Println("Hello ")

}
