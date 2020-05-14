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
    writeErr := ioutil.WriteFile(
        strings.Join([]string{homeDir,".vimrc"},"/"),
        vimConfigFile,
        os.Filemode(0777),
    )
    if writeErr != nil {
        fmt.Println(writeErr)
        return
    }
}
