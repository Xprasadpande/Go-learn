package main

import "fmt"

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

func main() {
	// fmt.Println(concat("Hello ", "World!"))
	// fmt.Println(concat("prasad ", "pande"))

	// fmt.Println(add(5, 7))

	sendsSoFar := 450
	const sendsToAdd = 50

	sendsSoFar = incrementSend(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

// func add(x, y int) int {
// 	return x + y
// }

func incrementSend(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
