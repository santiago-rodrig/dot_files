package main

import (
    "os"
    "os/exec"
    "log"
    "io/ioutil"
    "strings"
    "time"
)

func main() {
    // sets the user home directory
    homeDir := os.Getenv("HOME")
    workingDir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    // gets the contents of the vim configuration file
    vimConfigFile, err := ioutil.ReadFile(
        strings.Join([]string{workingDir,"/assets/vimrc"},"/"),
    )
    if err != nil {
        log.Fatal(err)
        return
    }
    // writes the contents to the vim configuration file
    err = ioutil.WriteFile(
        strings.Join([]string{homeDir,".vimrc"},"/"),
        vimConfigFile,
        os.FileMode(0777),
    )
    if err != nil {
        log.Fatal(err)
        return
    }
    time.Sleep(time.Second)
    log.Println("$HOME/.vimrc has been set up")
    // look for the existence of git on the system
    path, err := exec.LookPath("curl")
    if err != nil {
        log.Println("Please install curl: https://curl.haxx.se/download.html")
        log.Fatal(err)
    }
}
