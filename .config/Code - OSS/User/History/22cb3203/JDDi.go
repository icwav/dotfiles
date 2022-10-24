package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.OpenFile("testfile.txt",os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	
	data := ("hsfdsdf")

	for fileScanner.Scan() {
		if (fileScanner.Text() == "hello2" ) {
			_, err = f.WriteString(data)
			if err != nil{
				fmt.Println(err)
			}
		}
	readFile.Close()
	}
}