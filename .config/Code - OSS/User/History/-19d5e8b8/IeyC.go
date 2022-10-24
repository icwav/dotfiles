package main
// Token: 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
import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"

)

func main() {

	newetc := "/home/ivan/CCDC/testdir/newconfigs/Telnet"
	oldetc := "/home/ivan/CCDC/testdir/oldconfigs/telnet"
	oetc := "/home/ivan/CCDC/OLD_TELNET"

	
	bytesRead, err := ioutil.ReadFile(oetc)
	if err != nil{
		log.Fatal(err)
	}

	err = ioutil.WriteFile(oldetc, bytesRead, 0644)
	if err != nil {
		log.Fatal(err)
	}

	bytesRead1, err1 := ioutil.ReadFile(newetc)
	if err != nil {
		log.Fatal(err1)
	}
	fmt.Println("Configuring telnet...")
	err1 = ioutil.WriteFile(oetc, bytesRead1, 0644)
	if err != nil {
		log.Fatal(err1)
	} else {
		fmt.Println("telnet configured.")
	}
}