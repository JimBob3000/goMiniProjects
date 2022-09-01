package main

import (
	"errors"
	"fmt"
)

type node struct {
	data       string
	nextNodeID int
}

func main() {
	list := []node{}
	list = append(list, node{"hello ", 4})
	list = append(list, node{"James.", 2})
	list = append(list, node{",", 3})
	list = append(list, node{"it's ", 1})
	list = append(list, node{"world", 2})

	err := isValidLinkedList(&list)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Linked list is valid")
	}
}

func isValidLinkedList(list *[]node) error {
	visitedNodes := []int{}
	currentIndex := 0

	for {
		visitedNodes = append(visitedNodes, currentIndex)

		if (*list)[currentIndex].nextNodeID < 0 || (*list)[currentIndex].nextNodeID > len(*list) {
			return errors.New("invalid linked list, next node does not exist")
		} else if (*list)[currentIndex].nextNodeID != 0 {
			alreadyVisited := inNodeList(&visitedNodes, (*list)[currentIndex].nextNodeID)

			if alreadyVisited {
				return errors.New("invalid linked list, looping chain")
			}

			currentIndex = (*list)[currentIndex].nextNodeID
		} else {
			break
		}
	}

	if len(visitedNodes) != len(*list) {
		return errors.New("invalid linked list, not all nodes are visited")
	}

	return nil
}

func inNodeList(nodeList *[]int, nodeIndex int) bool {
	for i := 0; i < len(*nodeList); i++ {
		if (*nodeList)[i] == nodeIndex {
			return true
		}
	}

	return false
}
