package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "time"
)

func main() {
    // sets the user home directory
    homeDir := os.Getenv("HOME")
    workingDir, err := os.Getwd()
    if err != nil {
	    fmt.Println(err)
    }
    // gets the contents of the vim configuration file
    vimConfigFile, err := ioutil.ReadFile(
	    strings.Join([]string{workingDir,"/assets/vimrc"},"/"),
    )
    if err != nil {
        fmt.Println(err)
        return
    }
    // writes the contents to the vim configuration file
    err := ioutil.WriteFile(
        strings.Join([]string{homeDir,".vimrc"},"/"),
        vimConfigFile,
        os.FileMode(0777),
    )
    if err != nil {
        fmt.Println(writeErr)
        return
    }
    time.Sleep(time.Second)
    fmt.Println("$HOME/.vimrc has been set up")
}
