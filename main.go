package main

import (
    "strings"
    "github.com/santiago-rodrig/vim-setup/download"
    "os/exec"
    "log"
)

const colorReset string = "\033[0m"
const colorRed string = "\033[31m"
const colorGreen string = "\033[32m"

func installVimPlugins() (err error) {
    cmd := exec.Command(
        "vim",
        "+PlugInstall",
        "+qa",
    )
    err = cmd.Run()
    return
}

func buildSuccessMessage(msg string) (successMessage string) {
    successMessage := strings.Join(
        []string{
            colorGreen,
            "Success",
            colorReset,
            ":",
            " ",
            msg,
        },
        "",
    )
    return
}

func buildErrorMessage(msg string) (errorMessage string) {
    errorMessage := strings.Join(
        []string{
            colorRed,
            "Error",
            colorReset,
            ":",
            " ",
            msg,
        },
        "",
    )
    return
}

func checkDependencies() (msg string, err error) {
    _, err = exec.LookPath("vim")
    if err != nil {
        msg = buildErrorMessage("vim is not present in the PATH")
        return
    }
    _, err = exec.LookPath("git")
    if err != nil {
        msg = buildErrorMessage("git is not present in the PATH")
    }
    msg = buildSuccessMessage("this system has all the required dependencies")
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
        msg = buildErrorMessage(
            "vim plugins manager could not be set up",
        )
        log.Fatal(msg)
    } else {
        msg = buildSuccessMessage(
            "vim plugins manager has been set up",
        )
        log.Println(msg)
    }
    err = download.FetchVimrc()
    if err != nil {
        msg = buildErrorMessage(
            "vim configuration file could not be set up",
        )
        log.Fatal(msg)
    } else {
        msg = buildSuccessMessage(
            "vim configuration file has been set up",
        )
        log.Println(msg)
    }
    // install vim plugins
    err = installVimPlugins()
    if err != nil {
        msg = buildErrorMessage(
            "installing plugins with vim plug did not work",
        )
        log.Fatal(msg)
    }
    msg = buildSuccessMessage("vim plugins successfully installed")
    log.Println(msg)
}
