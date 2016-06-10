PassGen
=======
PassGen is a deterministic password generator which is generate same password for same inputs. The PassGen secured by [One Time Password](https://tools.ietf.org/html/rfc6238) (OTP) which is compatible with [Google Authenticator App](https://support.google.com/accounts/answer/1066447?hl=en).

Setup
--------
    passgen --init [--keyFile=PATH_TO_KEYFILE]
Initial setup will generate an encrypted file which is contain some configuration. Default key file path is `$HOME/.passgenkey`.  You need grant a user name and the password length (minimum value: 8) and some password parameter names. These parameters will be used for generating password. Leave blank the last `Password parameter` if you do not want to add more parameter.

    $ passgen --init
    Start PassGen Setup
    Your email address or user name or nick or whaterver: dzsubek
    Passwords length: 16
    Password parameter: Password for
    Password parameter: First dog name
    Password parameter: Favorite car
    Password parameter: Favorite number
    Password parameter:
    Your settings saved to: /Users/dzsubek/.passgenkey
    
    Scan this code with Authenticator App
    IN THIS PLACE YOU WILL SEE A QRCODE, AND YOU COULD SCAN IT WITH GOOGLE AUTHENTICATOR APP

> Every initialization will generate a new secret salt for password generation and you will get different passwords for the same input. Keep you `passgenkey` file in a secret place.

> Be careful! Do not use too easy parameters like these!

Usage
------
    passgen [--keyFile=PATH_TO_KEYFILE]
You need enter the current Passcode (OTP) form the Authenticator app then you need fill the parameter values and you will get the password.

    $ passgen
    Passcode: 687203
    Password for: GitHub
    First dog name: Csurka
    Favorite car: Trabant
    Favorite number: 13
    YmUzOGZjYWI1MTlh