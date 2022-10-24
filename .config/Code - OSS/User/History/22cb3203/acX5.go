package main

import (
	"bufio"
	"fmt"
	"os"
)

str= "hello"

func main() {
	readFile, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if (fileScanner.Text() !== str ) {
			fmt.Println("hello")
		}
	readFile.Close()
	}
}