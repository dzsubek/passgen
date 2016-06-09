package main

import (
    "bufio"
    "os"
    "os/user"
    "strings"
    "fmt"
    "path/filepath"
)

func readParam(name string) (value string) {
    fmt.Print(name)
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if (err != nil) {
        panic(err)
    }
    text = strings.TrimRight(text, "\r\n")

    return text
}

func getUserHome() (home string) {
    usr, err := user.Current()
    if err != nil {
        panic(err)
    }

    return usr.HomeDir + fmt.Sprintf("%c", filepath.Separator)
}
