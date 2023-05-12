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
	secretNum := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNum)

	fmt.Println("Please input your guess: ")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input, please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\r\n")
		// a,_:=strconv.Atoi(input)
		// fmt.Println(input);
		guess, err := strconv.Atoi(input)
		// fmt.Println(guess);
		if err != nil {
			fmt.Println("invalid input , please input an integer")
			continue
		}
		fmt.Println("your guess is ", guess)
		if guess > maxNum {
			fmt.Println("your guess is bigger than the secretNum. Please try again")
		} else if guess < maxNum {
			fmt.Println("your guess is smaller than the secretNum. Please try again")
		} else {
			fmt.Println("Congratulation! you are right!")
			break
		}
	}

}
