package main

import (
	"belajar_native/configs"
	"belajar_native/models"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	con, err := configs.ConnectRabbitMq()
	if err != nil {
		log.Fatal(err)
	}

	channel, err := con.Channel()
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := channel.Consume(
		"mitra_tani", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	db, err := configs.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for message := range msgs {
			// Unmarshal the payload into a struct
			var data models.SptaConsumer
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Error decoding message: %s\n", err)
				continue
			}

			log.Printf("Received message:\n%+v\n", data)
			opts := options.Update().SetUpsert(true)
			filter := bson.D{{"id", data.ID}}
			update := bson.D{{"$set", data}}

			_, err = db.Database("mitra_tani_consumer").Collection("spta_consumer").UpdateMany(nil, filter, update, opts)
			if err != nil {
				log.Printf("Error updating document: %s\n", err)
				continue
			}

		}
		defer con.Close()
		defer channel.Close()
		defer db.Disconnect(context.TODO())
	}()

	<-forever

}
