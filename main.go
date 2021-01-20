package main

import (
        "fmt"
        "log"
        "io/ioutil"
        "net/http"
        "os"
        "flag"
        )

func main() {
        //print commandline Args
        //fmt.Println(os.Args)

        //argLength := len(os.Args[1:])
        for i, a := range os.Args[1:] {
                fmt.Printf("Arg %d is %s\n", i+1, a)
        }

        var token string
        var group string

        flag.StringVar(&token, "t", "", "Specify Access Token.")
        flag.StringVar(&group, "g", "group", "Specify GroupID")
        //flag.StringVar(&token, "m", "messages", "Get Messages from group")
        flag.Parse()

        token   := ""
        group   := ""
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

