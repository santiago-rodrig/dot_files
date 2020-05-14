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
    homeDir := os.GetEnv("HOME")
    // gets the contents of the vim configuration file
    vimConfigFile, err := ioutil.ReadFile("./assets/vimrc")
    if err != nil {
        fmt.Println(err)
        return
    }
    // writes the contents to the vim configuration file
    writeErr := ioutil.WriteFile(
        strings.Join([]string{homeDir,".vimrc"},"/"),
        vimConfigFile,
        os.Filemode(0777),
    )
    if writeErr != nil {
        fmt.Println(writeErr)
        return
    }
    time.Sleep(time.Second)
    fmt.Println("$HOME/.vimrc has been set up")
}
