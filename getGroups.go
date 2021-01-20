package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	//"os"
	"flag"
	)

func main() {
	//print commandline Args
	//fmt.Println(os.Args)

	//argLength := len(os.Args[1:])
	//for i, a := range os.Args[1:] {
		//fmt.Printf("Arg %d is %s\n", i+1, a)
	//}

	var token string
	//var group string

	flag.StringVar(&token, "t", "", "Specify token.")
	//flag.StringVar(&pass, "p", "password", "Specify password. Default is password")
	flag.Parse()

	//token	:= "c0ca4f90335b0139de570e5f36a38d0b"
	//group	:= "65225076"

	resp, err  := http.Get("https://api.groupme.com/v3/groups?token=" + token)

	msgBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", msgBody)
	if err != nil {
		log.Fatal(err)
	}
}
