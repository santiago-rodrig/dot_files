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
    os.Remove(strings.Join([]string{homeDir,".vimrc"},"/"))
    // creates a new .vimrc file
    userVimConfigFile, createErr := os.Create(
        strings.Join([]string{homeDir,".vimrc"},"/"),
    )
    defer userVimConfigFile.Close()
    if createErr != nil {
        fmt.Println(createErr)
        return
    }
}
