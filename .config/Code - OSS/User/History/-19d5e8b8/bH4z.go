package main
// Token: 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
import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"os"

)

func main() {

	//newetc := "/home/ivan/CCDC/testdir/newconfigs/Telnet"
	oldetc := "/home/ivan/CCDC/testdir/oldconfigs/telnet"
	oetc := "/home/ivan/CCDC/OLD_TELNET"

	trequest, err := http.Get("https://raw.githubusercontent.com/CSUSB-CISO/godaedalus/main/Configurations/telnet")
	if err != nil {
		log.Fatalln(err)
	}

	tbody, err := ioutil.ReadAll(trequest.Body)
	if err != nil {
		log.Fatalln(err)
	}

//	telnetfile := string(tbody)

	file, err := os.Create(string(tbody))
	if err != nil {
		log.Fatal(err)
	} else {
		file.WriteString(file)
		
	}

	bytesRead, err := ioutil.ReadFile(oetc)
	if err != nil{
		log.Fatal(err)
	}

	err = ioutil.WriteFile(oldetc, bytesRead, 0644)
	if err != nil {
		log.Fatal(err)
	}


	bytesRead1, err1 := os.ReadFile(new)
	if err != nil {
		log.Fatal(err1)
	}


	fmt.Println("Configuring telnet...")
	err1 = ioutil.WriteFile(new, bytesRead1, 0644)
	if err != nil {
		log.Fatal(err1)
	} else {
		fmt.Println("telnet configured.")
	}
	fmt.Println(err)
}