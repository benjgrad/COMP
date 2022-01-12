package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"

	"strconv"
	"log"

	"github.com/benjgrad/COMP/src/comp/datastructures/linkedlist"
	"github.com/benjgrad/COMP/src/comp/algorithms"
)

func main() {

	list := linkedlist.LList{}
	
	reader := bufio.NewReader(os.Stdin)

	var command string

	if (len(os.Args) < 2) {
		fmt.Println("Enter a command:\n\tlinked list\n\tbinary search")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		command = input
	} else {
		command = strings.ToLower(os.Args[1])
	}

	switch command {
	case "linked list":
		L: for {
			fmt.Println("Enter an integer to prepend; Enter 'print' to view list; Enter 'exit' to finish")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input = strings.TrimSpace(input)

			switch input {
			case "exit":
				break L
			case "print":
				list.Print()
			default:
				fmt.Println("prepending: "+ input)
				val, err := strconv.Atoi(input)
				if err != nil {
					log.Fatal(err)
				}
				list.Prepend(&linkedlist.Node{Value: val})
			}
		}
	case "binary search":
		fmt.Println("Enter an ordered, space-delimited list of integers")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		
		list := strings.Split(input, " ")
		var intList = make([]int, len(list))
    
		for idx, i := range list {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			intList[idx] = j
		}

		fmt.Println("Enter a target value")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		target, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		fmt.Println("The index of the target is: ", algorithms.BinarySearch(intList, target))
		
	default:
		fmt.Println("Invalid argument: "+ command)
	}

}
