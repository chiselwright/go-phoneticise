package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phoneticise"
	"strings"
)

func main() {
	val := phoneticise.Phoneticise("Charlie")
	fmt.Println(val)

	val = phoneticise.Phoneticise("one two")
	fmt.Println(val)

	for {
		var input string

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("unable to read input")
		}

		input = strings.TrimSuffix(input, "\n") // remove newline

		fmt.Println("» " + phoneticise.Phoneticise(input))
	}
}
