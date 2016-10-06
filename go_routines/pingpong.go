package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player(table, "one")
	go player(table, "two")
	go player(table, "three")

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int, s string) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
		fmt.Println(ball, s)
	}
}
