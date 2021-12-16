package main

import (
	"context"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	/*topic := "test"
	fmt.Println(util.MACONSTANTE)
	util.AfficherUnTruc()
	client := connect("tcp://localhost:1883", "go_mqtt_client2")
	//token := client.Subscribe(topic, 1, nil)
	//token.Wait()
	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("Ma publication %d", i)
		client.Publish(topic, 2, false, text)
		time.Sleep(time.Second)
	}*/

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Airport:Airport@cluster0.0c6je.mongodb.net/AirportDataBase?retryWrites=true&w=majority"))
	err = client.Ping(ctx, readpref.Primary())
	fmt.Sprintf("%d", err)

	collection := client.Database("AirportDataBase").Collection("Pressure")
	res, err := collection.InsertOne(ctx, bson.D{{"idCaptor", 1}, {"iATA", "TLS"}, {"value", 40}, {"pickingDate", "2021-12-12T19:51:02.285+00:00"}})
	id := res.InsertedID
	fmt.Sprintf("%d", id)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

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
