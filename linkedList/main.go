package main

import "fmt"

type node struct {
	data       string
	nextNodeID int
}

func main() {
	list := []node{}
	list = append(list, node{"hello ", 4})
	list = append(list, node{"James.", 0})
	list = append(list, node{",", 3})
	list = append(list, node{"it's ", 1})
	list = append(list, node{"world", 2})

	currentIndex := 0
	fmt.Printf("STARTING LINKED LIST AT INDEX %v\n", currentIndex)

	for {
		fmt.Printf("Node id %v data: %v\n", currentIndex, list[currentIndex].data)
		if list[currentIndex].nextNodeID != 0 {
			currentIndex = list[currentIndex].nextNodeID
		} else {
			fmt.Println("END OF LINKED LIST")
			break
		}
	}
}
