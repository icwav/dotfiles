package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"os"
)


func Telnet() {
/*	err := os.Create("telnet")
	if err != nil {
		log.Fatal(err)
	}
*/
telurl := ("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")

	telrequest, err := http.Get(telurl)
	if err != nil {
		log.Fatalln(err)
	}

	telbody, err := ioutil.ReadAll(telrequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file2write := []byte(string(telbody))

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