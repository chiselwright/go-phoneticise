package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chiselwright/go-phoneticise"
)

func main() {
	for {
		var input string

		fmt.Println("{empty input or ^C to quit}")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("unable to read input")
		}

		input = strings.TrimSuffix(input, "\n") // remove newline

		if len(input) == 0 {
			break
		}

		fmt.Println("» " + phoneticise.Phoneticise(input))
	}
}
