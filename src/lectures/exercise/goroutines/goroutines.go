//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var sumTotal []int

func readAndSumFile(filename string) {

	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Getwd", err)
	}
	//fmt.Println("path->", path)

	filePath := path + `\exercise\goroutines\` + filename

	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open->", err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var sum int
	for scanner.Scan() {
		tmpint, err := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Printf(".")
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(filename + "-> nem lehet konvertalni")
			fmt.Println(scanner.Text() + " continue")
			continue
		}
		sum += int(tmpint)
	}
	fmt.Printf("vege\n")
	//return sum
	sumTotal = append(sumTotal, sum)
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	//var sumTotal []int

	for _, filename := range files {
		//filesum := readAndSumFile(filename)
		go readAndSumFile(filename)
		//fmt.Printf("filenum %v filesum-> %v\n", filenum, filesum)
		//sumTotal = append(sumTotal, filesum)

	}

	time.Sleep(1000 * time.Millisecond)
	var grandSum int

	for i := 0; i < len(sumTotal); i++ {
		grandSum += sumTotal[i]
		fmt.Println("sumTotal[", i, "]->", sumTotal[i])
	}
	fmt.Println("grandSum->", grandSum)
}
