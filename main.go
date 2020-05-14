package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    // sets the user home directory
    homeDir := os.GetEnv("HOME")
    // gets the contents of the vim configuration file
    vimConfigFile, err := ioutil.ReadFile("./assets/vimrc")
    if err != nil {
        fmt.Println(err)
        return
    }
    // deletes existing user vimrc if there is any
    err := os.Remove(strings.Join([]string{homeDir,".vimrc"},"/"))
    if err != nil {
        fmt.Println(err)
    }
}
