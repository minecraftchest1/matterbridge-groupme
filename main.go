package main

import (
        "fmt"
        "log"
        "io/ioutil"
        "net/http"
        )

func main() {
    //fmt.Println("Hello World")

        token   := "<token>"
        group   := "<group>"

        resp, err  := http.Get("https://api.groupme.com/v3/groups/" + group + "?token=" + token)

        msgBody, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        fmt.Printf("%s", msgBody)
        if err != nil {
                log.Fatal(err)
        }
}
