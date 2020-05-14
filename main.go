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
    if err != nil {
        log.Fatal("Error: Installing plugins with vim plug did not work")
        return
    }
    log.Println("Vim plugins successfully installed")
    return
}

func main() {
    // make the preparations
    download.FetchVimPlug()
    download.FetchVimrc()
    // install vim plugins
    installVimPlugins()
}
