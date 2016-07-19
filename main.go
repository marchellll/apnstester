package main

import (
  apns "github.com/sideshow/apns2"
  "github.com/sideshow/apns2/certificate"
  "github.com/davecgh/go-spew/spew"
  "github.com/spf13/viper"
  "log"
  "os"
  "fmt"
)

type Conf struct{
  P12filePath string
  P12password string
  Topic string
}

func getConfig() (*Conf) {
  viper.SetDefault("P12filePath", "./apns_sit2.p12")
  viper.SetDefault("P12password", "P@ssw0rd")
  viper.SetDefault("Topic", "com.example.Apps")


  viper.SetConfigName("config")
  viper.AddConfigPath(".")


  
  err := viper.ReadInConfig() // Find and read the config file
  if err != nil { // Handle errors reading the config file
      panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }
 
  conf := new(Conf)
  conf.P12filePath = viper.Get("P12filePath").(string)
  conf.P12password = viper.Get("P12password").(string)
  conf.Topic = viper.Get("Topic").(string)

  log.Println("Config :  ")
  spew.Dump(conf)
  log.Println("+++++++++++++++++++++")
  log.Println("+++++++++++++++++++++")
  log.Println("+++++++++++++++++++++")
  log.Println("+++++++++++++++++++++")

  return conf

}

func main() {

  conf := getConfig();

  

  cert, pemErr := certificate.FromP12File(conf.P12filePath , conf.P12password)
  if pemErr != nil {
    log.Println("Cert Error:", pemErr)
  }

  args := os.Args
  payloadString := "Lalala";
  if len(args) > 1 {
    payloadString = args[1]
  }


  notification := &apns.Notification{}
  notification.DeviceToken = "XXXXXXXXX54FD9824E88A458A225DD31C79EC4CEFD0C6DFE46C0123456789012"
  notification.Topic = conf.Topic
  
  notification.Payload = []byte(fmt.Sprintf(`{"aps":{"alert":"%s","sound":"default"},"data":[]}`, payloadString)) // See Payload section below

  if len(args) > 2 {
    notification.DeviceToken = args[2]
  }

  log.Println("Pushing ")
  spew.Dump(notification)

  client := apns.NewClient(cert).Production()
  res, err := client.Push(notification)


  if err != nil {
    log.Println("Error:", err)
    return
  }

  log.Println("APNs ID:", res.ApnsID)
}