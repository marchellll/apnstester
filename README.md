# apnstester
Simple tester for APNS connection. I use this to test if an server is able to trigger an apple notification.

I depend on these work of other people to create this :

  - "github.com/sideshow/apns2"
  - "github.com/sideshow/apns2/certificate"
  - "github.com/davecgh/go-spew/spew"
  - "github.com/spf13/viper"

# Preparation
Please edit config.json as needed. And ofcourse you need your **.p12** file.

```json 
{
  "P12filePath" : "./apns.p12", 
  "P12password": "P@ssw0rd",
  "Topic": "com.example.Apps"
}
```

```P12filePath``` is your where .p12 file. 
```P12password``` is your .p12 password. 
```Topic``` is your App ID. 

# Usage
build it. then run it.
**1st** argument is The payload,
**2nd** argument is your device token

```sh
go build
./testAPNS "Yooohhooooooo" XXXXXXXXXXXXD9824E88A458A225DD31C79EC4CEFD0C6DFE46C03AAFF3A4E123123
```
