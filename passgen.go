package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "encoding/json"
    "github.com/dgryski/dgoogauth"
)

var secret = "YPC3RyUa4XemGPhQ"

var setup = flag.Bool("init", false, "initialize new keys")
var keyFile = flag.String("keyFile", "", "initialize new keys")

type KeyInfo struct {
    Name string
    Otp dgoogauth.OTPConfig
    Params []string
    PassLength int
}

func init() {
    flag.Parse()
    if (*keyFile == "") {
        *keyFile = getUserHome() + ".passgenkey"
    }
}

func getInfo() KeyInfo {
    file, err := ioutil.ReadFile(*keyFile)
    if (err != nil) {
        fmt.Println(err)
        panic(err)
    }
    
    var params KeyInfo
    json.Unmarshal(decrypt(file), &params)

    return params
}

func main() {
    if *setup {
        initialize()
    } else {
        info := getInfo()
        otp := info.Otp
        var passcode = readParam("Passcode: ")
        
        success, err := otp.Authenticate(passcode)
        if(!success || err != nil) {
            panic("Invalid passcode")
        }

        var params string
        for _, param := range info.Params {
            params = params + readParam(param + ": ")
        }
        
        pass := createPass(params, info.Otp.Secret)
        for len(pass) < info.PassLength {
            pass = pass + createPass(params, pass)
        }
        pass = pass[:info.PassLength]
        fmt.Println(pass)
    }
}
