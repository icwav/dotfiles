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
		fmt.Println(fileScanner.Text())
		if fileScanner.Scan() = ("All the data I wish to write to a file")
		os.WriteFile("testfile.txt"), []byte("helloo"), 0666; err
	}

	readFile.Close()
}r