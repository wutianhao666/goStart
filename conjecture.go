package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input your guess:")
		input, err := reader.ReadString('\n')
		//fmt.Println("what i read :", input)
		if err != nil {
			fmt.Println("An error occured while reading input.", err)
			continue
		}
		input = strings.TrimSuffix(input, "\r\n")
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invaild input2222.", "guess:", guess, err)
			continue
		}
		fmt.Println("You guess is ", guess)

		if guess > secretNumber {
			fmt.Println("You guess is bigger than answer!")
		} else if guess < secretNumber {
			fmt.Println("You guess is smaller than answer!")
		} else {
			fmt.Println("You guess is correct!")
			break
		}
	}
}
