package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helping functions and types.

type conditionCallback func() bool
type resetCallback func()
type body func(index int)

func while(index int, condition conditionCallback, conditionReset resetCallback, loopBody body) {
	conditionResult := condition()
	if conditionResult {
		loopBody(index)
		conditionReset()
		while(index+1, condition, conditionReset, loopBody)
	}
}

func reduceValue(value int, loopBody body) {
	while(0, func() bool {
		return value > 0
	}, func() {
		value--
	}, loopBody)
}

func printAllValues(results []int) {
	size := len(results)
	reduceValue(size, func(index int) {
		fmt.Println(results[index])
	})
}

// Logic

func runTestCase() int {
	var amountOfNumbersToBeRead int
	fmt.Scan(&amountOfNumbersToBeRead)

	reader := bufio.NewReader(os.Stdin)
	numbersString, _ := reader.ReadString('\n')

	numbers := strings.Fields(numbersString)

	finalNumber := 0

	reduceValue(amountOfNumbersToBeRead, func(index int) {
		currentNumber, _ := strconv.Atoi(numbers[index])
		if currentNumber > 0 {
			finalNumber += (currentNumber * currentNumber)
		}
	})

	return finalNumber
}

func main() {
	var numberOfTests int
	_, error := fmt.Scanln(&numberOfTests)

	if error != nil {
		fmt.Println(error)
	}

	var results []int
	reduceValue(numberOfTests, func(index int) {
		results = append(results, runTestCase())
	})

	printAllValues(results)
}
