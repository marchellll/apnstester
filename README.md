# apnstester
Simple tester for APNS connection. I use this to test if a server is able to trigger an apple notification.

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

```P12filePath``` is your .p12 file path. 
```P12password``` is your .p12 password. 
```Topic``` is your App ID. 

# Build

You need Go installed in your system to build it. https://golang.org/

Simple usage would be run directly : 
```sh
go run main.go "Yooohhooooooo" XXXXXXXXXXXXD9824E88A458A225DD31C79EC4CEFD0C6DFE46C03AAFF3A4E123123
```
**1st** argument is The payload,
**2nd** argument is your device token

# Runnables

There is convenient prebuilded executables generated in OSX. You could use these if you don't want to build it first.

```sh
./apnstester-darwin "Yeyeyeye" XXXXXXXXXXXXD9824E88A458A225DD31C79EC4CEFD0C6DFE46C03AAFF3A4E123123
```

```apnstester-darwin``` generated using ```go build```

```apnstester-linux``` generated using ```CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apnstester-linux .```
