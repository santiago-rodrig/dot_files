package download

import (
    "os"
    "net/http"
    "io"
    "log"
)

func handleError(err error) {
    if err != nil {
        log.Fatal("Error:", err)
    }
}

func prepareVimFolder() (err error) {
    fileMode := os.FileMode(os.ModeDir)
    home, err := os.UserHomeDir()
    handleError(err)
    err = os.MkdirAll(home + "/.vim/autoload", fileMode)
    handleError(err)
    return
}

func download(url string) (err error) {
}

// Gets the contents of the vimrc file under assets
// and writes them to $HOME/.vimrc
func FetchVimrc() error {
}

// Gets the contents of plug.vim and writes them to
// $HOME/.vim/autoload/plug.vim
func FetchVimPlug() error {
}
