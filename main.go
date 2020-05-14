package main

import (
    "os"
    "fmt"
    "io/ioutil"
)

var homeDir string
var vimConfigFile os.File

func main() {
    // sets the user home directory
    homeDir = os.GetEnv("HOME")
    // gets the contents of the vim configuration file
    vimConfigFile, err = ioutil.ReadFile("./assets/vimrc")
    if err != nil {
        fmt.Println(err)
        return
    }
}
