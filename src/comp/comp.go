package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"

	"strconv"
	"log"

	"github.com/benjgrad/COMP/src/comp/datastructures/linkedlist"
)

func main() {

	list := linkedlist.LList{}
	
	reader := bufio.NewReader(os.Stdin)

	switch command := strings.ToLower(os.Args[1]); command {
	case "linked list":
		L: for {
			fmt.Println("Enter number to prepend; Enter 'print' to view list; Enter 'exit' to finish")
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
	default:
		fmt.Println("Invalid argument: "+ command)
	}

}
