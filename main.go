package main

import (
	"time"
	"os"
	"strconv"
	"fmt"
)

func main() {
	arg := os.Args

	if len(arg) == 1 {
		fmt.Println("Input number")
		return
	}
	num, err := strconv.Atoi(arg[1])

	if err != nil {
		fmt.Println("Not a number")
		return
	}

	sleep(num)
}

func sleep(number int) {
	time.Sleep(time.Second * time.Duration(number))
}

//func findBigPrime(array []int) (int, error) {
//	max := -1
//
//	for _, element := range array {
//		if isPrime(element) && element > max {
//			max = element
//		}
//	}
//
//	if max == -1 {
//		return max, errors.New("Простого числа нет")
//	}
//
//	return max, nil
//}
//
//func isPrime(number int) bool {
//	if number == 2 {
//		return true
//	}
//
//	if number%2 == 0 || number < 2 {
//		return false
//	}
//
//	for i := 3; i < int(math.Sqrt(float64(number))); i += 2 {
//		if number%i == 0 {
//			return false
//		}
//	}
//	return true
//}
