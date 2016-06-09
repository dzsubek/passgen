package main

import (
    "encoding/base32"
    "crypto"
    "github.com/dgryski/dgoogauth"
    "fmt"
    "crypto/rand"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "net/url"
)

func createOTP() (otp dgoogauth.OTPConfig){
    keySize := crypto.SHA1.Size()
	key := make([]byte, keySize)
	_, err := rand.Read(key)

    if (err != nil) {
        panic(err)
    }

    dgoogauth := dgoogauth.OTPConfig{}
    dgoogauth.Secret = base32.StdEncoding.EncodeToString(key)
    dgoogauth.WindowSize = 3
    dgoogauth.HotpCounter = 0
    dgoogauth.UTC = true
    
    return dgoogauth
}

func initialize() {
    fmt.Println("Start PassGen Setup")
    
    info := KeyInfo{}
    info.Name = readParam("Your email address or user name or nick or whaterver: ")
    len, err := strconv.Atoi(readParam("Passwords length: "))
    if (err != nil || len < 8) {
        panic("Invalid password length (min: 8)")
    }
    info.PassLength = len

    var params []string
    var text = readParam("Password parameter: ")
    for text != "" {
        params = append(params, text)
        text = readParam("Password parameter: ")
    }

    info.Params = params
    otp := createOTP()
    info.Otp = otp
    
    json, err := json.Marshal(info)
    if err != nil {
        panic(err)
    }
    
    err = ioutil.WriteFile(*keyFile, encrypt(json), 0600)
    if err != nil {
        panic(err)
    }

    fmt.Println("Your settings saved to: " + *keyFile)
    link := otp.ProvisionURIWithIssuer(url.QueryEscape(info.Name), "PassGen")
    printQR(link)
}