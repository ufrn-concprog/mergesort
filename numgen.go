// Generating a set of random numbers
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 3, 2023

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var factor float64
	var filename string
	if len(os.Args) == 3 {
		factor, _ = strconv.ParseFloat(os.Args[1], 64)
		filename = os.Args[2]
	} else {
		fmt.Println("Provided arguments are invalid")
		os.Exit(1)
	}

	filepath := "input/" + filename + ".in"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error when creating the file")
        os.Exit(1)
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := math.Pow(2, factor)
	for i := 0; i < int(numbers); i++ {
		number := random.Intn(int(numbers * 0.1) + 1)
		if _, err := file.Write([]byte(strconv.Itoa(number) + " ")); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	
	fmt.Println(filepath, "file generated.")
	file.Close()
}