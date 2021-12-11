package main

import (
	"fmt"
	"log"
	"time"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	topic := "test"
	fmt.Println(util.MACONSTANTE)
	util.AfficherUnTruc()
	client := connect("tcp://localhost:1883", "go_mqtt_client2")
	//token := client.Subscribe(topic, 1, nil)
	//token.Wait()
	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("Ma publication %d", i)
		client.Publish(topic, 2, false, text)
		time.Sleep(time.Second)
	}

}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetUsername(user)
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts

}

func connect(brokerURI string, clientId string) mqtt.Client {

	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client

}
