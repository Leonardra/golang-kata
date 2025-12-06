package main

import (
	"bufio"
	"fmt"
	"kata-one/doctor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var introPrompt string = doctor.DisplayIntroPrompt()
	fmt.Println(introPrompt)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		if userInput == "exit" {
			break
		} else {
			fmt.Printf("You entered: %s", userInput)
		}
	}

}
