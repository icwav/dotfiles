package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)




func main() {
	new, err := os.Create("telnet")
	if err != nil {
		log.Fatal(err)
	}

	trequest, err := http.Get("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	if err != nil {
		log.Fatalln(err)
	}

	tbody, err := ioutil.ReadAll(trequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file2write := []byte(string(tbody))

	what, err := os.Write(file2write)
	if err != nil {
		log.Fatalln(err)
	}


	final, err := os.WriteFile(new, what, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(final)

}