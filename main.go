package main

import (
    "github.com/santiago-rodrig/vim-setup/download"
    "os/exec"
    "log"
)

func installVimPlugins() (err error) {
    cmd := exec.Command(
        "vim",
        "+PlugInstall",
        "+qa",
    )
    err = cmd.Run()
    return
}

func checkDependencies() (msg string, err error) {
    _, err = exec.LookPath("vim")
    if err != nil {
        msg = "Error: vim is not present in the PATH"
        return
    }
    _, err = exec.LookPath("git")
    if err != nil {
        msg = "Error: git is not present in the PATH"
    }
    msg = "This system has all required dependencies"
    return
}

func main() {
    // check dependencies
    msg, err := checkDependencies()
    if err != nil {
        log.Fatal(msg)
    } else {
        log.Println(msg)
    }
    // make the preparations
    err = download.FetchVimPlug()
    if err != nil {
        log.Fatal("Error: vim plugins manager could not be set up")
    } else {
        log.Println("Success: vim plugins manager has been set up")
    }
    err = download.FetchVimrc()
    if err != nil {
        log.Fatal("Error: vim configuration file could not be set up")
    } else {
        log.Println("Success: vim configuration file has been set up")
    }
    // install vim plugins
    err = installVimPlugins()
    if err != nil {
        log.Fatal("Error: Installing plugins with vim plug did not work")
    }
    log.Println("Vim plugins successfully installed")
}
