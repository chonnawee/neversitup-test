package main

import (
	"bufio"
	"fmt"
	"neversitup-test/internal/core/implements/services/counter"
	"neversitup-test/internal/core/implements/services/odd"
	"neversitup-test/internal/core/implements/services/permutations"
	"os"
	"strconv"
	"strings"
)

func getCountSmileysAnswer(reader *bufio.Reader) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input. Please enter a string.")
		return
	}
	input = strings.TrimSpace(input)
	var slice []string
	smileys := strings.Split(input, ",")
	slice = append(slice, smileys...)
	service := counter.NewCounterServiceImplement()
	fmt.Println("answer is ", service.CountSmileyInSlice(slice))
}

func getFindTheOddAnswer(reader *bufio.Reader) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input. Please enter a string.")
		return
	}
	input = strings.TrimSpace(input)
	numbersStr := strings.Split(input, ",")

	var slice []int
	for _, numberStr := range numbersStr {
		numberInt, _ := strconv.Atoi(numberStr)
		slice = append(slice, numberInt)
	}
	service := odd.NewOddServiceImplement()
	fmt.Println("answer is ", service.FindIntInSlice(slice))
}

func getPermutationsAnswer(reader *bufio.Reader) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input. Please enter a string.")
		return
	}
	input = strings.TrimSpace(input)
	service := permutations.NewPermutationsServiceImplement()
	fmt.Println("answer is ", service.Generate(input))
}

func getMessageFromChoice(choice int, reader *bufio.Reader) {
	switch choice {
	case 1:
		fmt.Println("Enter the string for which you want to generate permutations.:")
		getPermutationsAnswer(reader)
	case 2:
		fmt.Println("Enter the numbers for which you want to find the odd occurrence.")
		fmt.Println("For example 1,3,3,5,5,6,6 :")
		getFindTheOddAnswer(reader)
	case 3:
		fmt.Println("Enter the smiley face you want to count")
		fmt.Println("For example :),:[ :")
		getCountSmileysAnswer(reader)
	}
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("--------------------------")
		fmt.Println("Select a test:")
		fmt.Println("1.Permutations")
		fmt.Println("2.Find the odd int")
		fmt.Println("3.Count the smiley faces!")
		fmt.Println("--------------------------")
		fmt.Println("Enter choice:")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		getMessageFromChoice(choice, reader)
	}
}
