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

	userInput, _ := reader.ReadString('\n')
	fmt.Printf("You entered: %s", userInput)
}
