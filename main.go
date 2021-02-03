package main

import (
        "fmt"
        "log"
        "io/ioutil"
        "net/http"
        "os"
        "flag"
        )

var token string
var group string
var get string

func main() {
        //print commandline Args
        //fmt.Println(os.Args)

        //argLength := len(os.Args[1:])
        for i, a := range os.Args[1:] {
                fmt.Printf("Arg %d is %s\n", i+1, a)
        }



        flag.StringVar(&token, "t", "token", "Specify Access Token.")
        flag.StringVar(&group, "g", "group", "Specify GroupID")
	flag.StringVar(&get, "get", "false", "Get group info.")
        //flag.StringVar(&token, "m", "messages", "Get Messages from group")
        flag.Parse()

	fmt.Printf("%s", token)
	fmt.Printf("%s", group)

        //token   := ""
        //group   := ""

	//if get == "true"{
		getMessages()
	//}
}
func getMessages() {
        resp, err  := http.Get("https://api.groupme.com/v3/groups/" + group + "?token=" + token)

        msgBody, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        fmt.Printf("%s", msgBody)
        if err != nil {
                log.Fatal(err)
	}
}

