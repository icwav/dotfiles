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
		str = "hello"
		if (fileScanner.Text() == "" ) {
			err := ioutil.WriteFile("testfile.txt", str, 0777)
			if err != nil{
				fmt.Println("handle this")
			}
		}
	readFile.Close()
	}
}