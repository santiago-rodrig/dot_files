package download

import (
    "strings"
    "os"
    "net/http"
    "io"
    "log"
)

// handles errors
func handleError(err error) bool {
    if err != nil {
        log.Fatal("Error:", err)
        return false
    }

    return true
}

// prepares $HOME/.vim/autoload
func prepareVimFolder() (err error) {
    fileMode := os.FileMode(os.ModeDir)
    home, err := os.UserHomeDir()
    if !handleError(err) { return }
    err = os.MkdirAll(home + "/.vim/autoload", fileMode)
    if !handleError(err) { return }
    return
}

// downloads from the given url
func download(url string) (body io.ReadCloser, err error) {
    response, err := http.Get(url)
    if !handleError(err) { return }
    body = response.Body
    return
}

// Gets the contents of the vimrc file under assets
// and writes them to $HOME/.vimrc
func FetchVimrc() (err error) {
    body, err := download(strings.Join([]string{
        "https://raw.githubusercontent.com/santiago-rodrig",
        "vim-setup/master/assets/vimrc",
    }, "/")
    if !handleError(err) { return }
    home, err := os.UserHomeDir()
    if !handleError(err) { return }
    vimrc, err := os.Create(home + "/.vimrc")
    if !handleError(err) { return }
    _, err = io.Copy(vimrc, body)
    return
}

// Gets the contents of plug.vim and writes them to
// $HOME/.vim/autoload/plug.vim
func FetchVimPlug() (err error) {
    body, err := download(strings.Join([]string{
        "https://raw.githubusercontent.com/junegunn",
        "vim-plug/master/plug.vim",
    }, "/")
    if !handleError(err) { return }
    home, err := os.UserHomeDir()
    if !handleError(err) { return }
    err = prepareVimFolder()
    if !handleError(err) { return }
    plug, err := os.Create(home + "/.vim/autoload/plug.vim")
    if !handleError(err) { return }
    io.Copy(plug, body)
    return
}
