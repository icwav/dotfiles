package main

import (
	"log"
	"io/ioutil"
	"net/http"
//	"fmt"
	"os"
)


func Telnet() {
/*	err := os.Create("telnet")
	if err != nil {
		log.Fatal(err)
	}
*/
url := ("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")

	trequest, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	tbody, err := ioutil.ReadAll(trequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file2write := []byte(string(tbody))

	/*what, err := os.Write(file2write)
	if err != nil {
		log.Fatalln(err)
	}
*/

	os.WriteFile("telnet", []byte(file2write), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}