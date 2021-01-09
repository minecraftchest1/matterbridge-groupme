package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	)

func main() {
    //fmt.Println("Hello World")

	token	:= "<Groupme Token>"

	resp, err  := http.Get("https://api.groupme.com/v3/groups/65225076?token=" + token)

	msgBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", msgBody)
	fmt.Println(err)
}
