package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	readFile, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	
	data := []byte("hello")

	for fileScanner.Scan() {
		if (fileScanner.Text() == "new data that wasn't there originally" ) {
			_, err = fmt.FPrintln(f , data)
			if err != nil{
				fmt.Println("handle this")
			}
		}
	readFile.Close()
	}
}