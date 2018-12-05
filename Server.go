package main

import (
	"fmt"
	"net/http"
	"os"
)

var port = "8080"

func WriteFile(s string) {
	/* opens a file and append our string to it */
	f, err := os.OpenFile("statics/ceos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cheackError(err)
	defer f.Close()
	f.WriteString(s)
}

func cheackError(e error) {
	/* we want to follow DRY principle and check for errors in a function  */
	if e != nil {
		panic(e)
	}
}

func handleCEO(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()

	cheackError(err)

	CEO := req.Form.Get("CEO")
	company := req.Form.Get("company")

	s := fmt.Sprintf("{\"company\":\"%s\",\"CEO\":\"%s\"} \n", company, CEO)

	WriteFile(s)

	fmt.Fprint(w, s)

}

func main() {

	http.HandleFunc("/CEO", handleCEO)
	fmt.Println("Server is runnig at port " + port)

	http.ListenAndServe(":"+port, nil)

}
