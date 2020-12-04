package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num[] int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		if scanner.Text() == "" {
			fmt.Println("Scanner Failed")
			break
		}
		n, er := strconv.Atoi(scanner.Text())
		if er != nil {
			panic(er)
		}
		num = append(num, n)
	}

	fmt.Println( "Sum =" , checkSum2020(num) );
}


func checkSum2020(num []int) int {
	for i := 0; i < len(num); i++ {
		for j := i; j < len(num); j++ {
			for k := j; k < len(num); k++ {
				if num[i] + num[j] + num[k] == 2020 {
					return num[i] * num[j] * num[k]
				}
			}
		}
	}
	return -1
}

