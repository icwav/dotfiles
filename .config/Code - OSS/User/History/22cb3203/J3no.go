package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	readFile, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if (fileScanner.Text() == "All the data I wish to write to a file") {
			fmt.Println("hello")
		}
	readFile.Close()
	}
}