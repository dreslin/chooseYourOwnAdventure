package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(question string, list ...[]string) string {
	fmt.Sprintf(question)
	inputReader := bufio.NewReader(os.Stdin)
	input1, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input1 = strings.TrimSuffix(input1, "\n")
	if err != nil {

		return getInput(question)
	}
	return input1

}

func wrongInput() {
	fmt.Println()
	fmt.Println("Alas! Thy speech is unintelligible!")
	fmt.Println()

}

// func showP1() {
// 	fmt.Println(player1)
// }
