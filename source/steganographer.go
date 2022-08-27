package main

import (
	"fmt"
	"strings"
)

func readMessage() {

}

func writeMessage() {

}

func main() {
	var programMode string
	fmt.Print("Do you want to write or read a image message? ")
	fmt.Scanf("%s", &programMode)

	programMode = strings.ToLower(programMode)
	for programMode != "write" && programMode != "read" {
		fmt.Print("Please, enter write or read. ")
		fmt.Scanf("%s", &programMode)

		programMode = strings.ToLower(programMode)
	}

	if programMode == "read" {
		readMessage()
	} else {
		writeMessage()
	}
}
