package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "encoding/json"
    "github.com/dgryski/dgoogauth"
    "github.com/howeyc/gopass"
    //"crypto/sha1"
)

var isSetup = flag.Bool("init", false, "setup new keys")
var keyFile = flag.String("keyFile", "", "setup new keys")

type KeyInfo struct {
    Name string
    Otp dgoogauth.OTPConfig
}

func init() {
    flag.Parse()
    if (*keyFile == "") {
        *keyFile = getUserHome() + ".passgenkey"
    }
}

func getInfo(masterPass []byte) KeyInfo {
    file, err := ioutil.ReadFile(*keyFile)
    if (err != nil) {
        fmt.Println(err)
        panic(err)
    }
    
    var info KeyInfo
    err = json.Unmarshal(decrypt(file, masterPass), &info)
    if (err != nil) {
        panic("Can not read configuration, may master password is wrong")
    }
    return info
}

func getMasterPass() []byte  {
    fmt.Printf("Enter master password: ")
    masterPass, err := gopass.GetPasswd()
    if err != nil {
        panic(err)
    }

    return []byte(masterPass)
}

func main() {
    if *isSetup {
        setup()
    } else {
        masterPass := getMasterPass()
        info := getInfo(masterPass)

        fmt.Println(info.Name);
        otp := info.Otp
        var passcode = readParam("Authenticator passcode: ")

        success, err := otp.Authenticate(passcode)
        if(!success || err != nil) {
             panic("Invalid passcode")
        }

		param := readParam("Generate password for: ")
        length := len(masterPass) + len(param)
        if length < 8 {
            length = 8
        }

		pass := createPass(param, masterPass, length)
		fmt.Println(pass)
    }
}
