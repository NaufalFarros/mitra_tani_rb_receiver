package main

import (
	"belajar_native/configs"
	"belajar_native/models"
	"context"
	"encoding/json"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
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

	db, err := configs.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// go InsertDataFromRBMQ(ctx, db, channel, &wg)
	// go DefaultUpsert(db, channel, &wg)
	go MyFuncInsert(db, channel, &wg)
	wg.Wait()
}

// func InsertDataFromRBMQ(ctx context.Context, db *mongo.Client, channel *amqp.Channel, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	msgs_ari, err := channel.Consume(
// 		"mitra_tani_ari", // queue
// 		"",               // consumer
// 		true,             // auto-ack
// 		false,            // exclusive
// 		false,            // no-local
// 		false,            // no-wait
// 		nil,              // arguments
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_meja_tebu, err := channel.Consume(
// 		"mitra_tani_meja_tebu", // queue
// 		"",                     // consumer
// 		true,                   // auto-ack
// 		false,                  // exclusive
// 		false,                  // no-local
// 		false,                  // no-wait
// 		nil,                    // arguments
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_selektor, err := channel.Consume(
// 		"mitra_tani_selektor", // queue
// 		"",                    // consumer
// 		true,                  // auto-ack
// 		false,                 // exclusive
// 		false,                 // no-local
// 		false,                 // no-wait
// 		nil,                   // arguments
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_spta, err := channel.Consume(
// 		"mitra_tani_spta", // queue
// 		"",                // consumer
// 		true,              // auto-ack
// 		false,             // exclusive
// 		false,             // no-local
// 		false,             // no-wait
// 		nil,               // arguments
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_timbangan, err := channel.Consume(
// 		"mitra_tani_timbangan", // queue
// 		"",                     // consumer
// 		true,                   // auto-ack
// 		false,                  // exclusive
// 		false,                  // no-local
// 		false,                  // no-wait
// 		nil,                    // arguments
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	AriChan := make(chan models.TAri)
// 	MejaTebuChan := make(chan models.TMejaTebu)
// 	TSelectorChan := make(chan models.TSelektor)
// 	SptaChan := make(chan models.TSpta)
// 	TTimbanganChan := make(chan models.TTimbang)

// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case message := <-msgs_ari:
// 				// Unmarshal the payload into a struct
// 				var data models.TAri
// 				err := json.Unmarshal(message.Body, &data)
// 				if err != nil {
// 					log.Printf("Error decoding message: %s\n", err)
// 					continue
// 				}
// 				AriChan <- data
// 			case message := <-msg_meja_tebu:
// 				// Unmarshal the payload into a struct
// 				var data models.TMejaTebu
// 				err := json.Unmarshal(message.Body, &data)
// 				if err != nil {
// 					log.Printf("Error decoding message: %s\n", err)
// 					continue
// 				}
// 				MejaTebuChan <- data
// 			case message := <-msg_selektor:
// 				// Unmarshal the payload into a struct
// 				var data models.TSelektor
// 				err := json.Unmarshal(message.Body, &data)
// 				if err != nil {
// 					log.Printf("Error decoding message: %s\n", err)
// 					continue
// 				}
// 				TSelectorChan <- data
// 			case message := <-msg_spta:
// 				// Unmarshal the payload into a struct
// 				var data models.TSpta
// 				err := json.Unmarshal(message.Body, &data)
// 				if err != nil {
// 					log.Printf("Error decoding message: %s\n", err)
// 					continue
// 				}
// 				SptaChan <- data
// 			case message := <-msg_timbangan:
// 				// Unmarshal the payload into a struct
// 				var data models.TTimbang
// 				err := json.Unmarshal(message.Body, &data)
// 				if err != nil {
// 					log.Printf("Error decoding message: %s\n", err)
// 					continue
// 				}
// 				TTimbanganChan <- data
// 			}
// 		}
// 	}()

// 	// Insert Data into MongoDB
// 	insertData := func(collectionName string, data interface{}) {
// 		coll := db.Database("mitra_tani_consumer").Collection(collectionName)
// 		_, err := coll.InsertOne(ctx, data)
// 		if err != nil {
// 			log.Printf("Error inserting data into collection %s: %s\n", collectionName, err)
// 		}
// 	}

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case data := <-AriChan:
// 			insertData("t_ari_mtani", data)
// 		case data := <-MejaTebuChan:
// 			insertData("t_meja_tebu_mtani", data)
// 		case data := <-TSelectorChan:
// 			insertData("t_selektor_mtani", data)
// 		case data := <-SptaChan:
// 			insertData("t_spta_mtani", data)
// 		case data := <-TTimbanganChan:
// 			insertData("t_timbangan_mtani", data)
// 		}
// 	}
// }

// func DefaultUpsert(db *mongo.Client, channel *amqp.Channel, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	msgs_ari, err := channel.Consume(
// 		"mitra_tani_ari", // queue
// 		"",               // consumer
// 		true,             // auto-ack
// 		false,            // exclusive
// 		false,            // no-local
// 		false,            // no-wait
// 		nil,              // arguments
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_meja_tebu, err := channel.Consume(
// 		"mitra_tani_meja_tebu", // queue
// 		"",                     // consumer
// 		true,                   // auto-ack
// 		false,                  // exclusive
// 		false,                  // no-local
// 		false,                  // no-wait
// 		nil,                    // arguments
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_selektor, err := channel.Consume(
// 		"mitra_tani_selektor", // queue
// 		"",                    // consumer
// 		true,                  // auto-ack
// 		false,                 // exclusive
// 		false,                 // no-local
// 		false,                 // no-wait
// 		nil,                   // arguments

// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_spta, err := channel.Consume(
// 		"mitra_tani_spta", // queue
// 		"",                // consumer
// 		true,              // auto-ack
// 		false,             // exclusive
// 		false,             // no-local
// 		false,             // no-wait
// 		nil,               // arguments
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	msg_timbangan, err := channel.Consume(
// 		"mitra_tani_timbangan", // queue
// 		"",                     // consumer
// 		true,                   // auto-ack
// 		false,                  // exclusive
// 		false,                  // no-local
// 		false,                  // no-wait
// 		nil,                    // arguments
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	go func() {
// 		for message := range msgs_ari {
// 			// Unmarshal the payload into a struct
// 			var data models.TAri
// 			err := json.Unmarshal(message.Body, &data)
// 			if err != nil {
// 				log.Printf("Error decoding message: %s\n", err)
// 				continue
// 			}

// 			log.Printf("Received message TAri :\n%+v\n", data)
// 			opts := options.Update().SetUpsert(true)
// 			filter := bson.D{{Key: "id", Value: data.IDAri}}
// 			update := bson.D{{Key: "$set", Value: data}}

// 			_, err = db.Database("mitra_tani_consumer").Collection("t_ari_mtani").UpdateMany(context.TODO(), filter, update, opts)
// 			if err != nil {
// 				log.Printf("Error updating document: %s\n", err)
// 				continue
// 			}

// 		}
// 		defer channel.Close()
// 		defer db.Disconnect(context.TODO())
// 	}()

// 	go func() {
// 		for message := range msg_meja_tebu {
// 			// Unmarshal the payload into a struct
// 			var data models.TMejaTebu
// 			err := json.Unmarshal(message.Body, &data)
// 			if err != nil {
// 				log.Printf("Error decoding message: %s\n", err)
// 				continue
// 			}

// 			log.Printf("Received message TMejaTebu:\n%+v\n", data)
// 			opts := options.Update().SetUpsert(true)
// 			filter := bson.D{{Key: "id", Value: data.IDMejatebu}}
// 			update := bson.D{{Key: "$set", Value: data}}

// 			_, err = db.Database("mitra_tani_consumer").Collection("t_meja_tebu_mtani").UpdateMany(context.TODO(), filter, update, opts)
// 			if err != nil {
// 				log.Printf("Error updating document: %s\n", err)
// 				continue
// 			}

// 		}

// 		defer channel.Close()
// 		defer db.Disconnect(context.TODO())
// 	}()

// 	go func() {
// 		for message := range msg_selektor {
// 			// Unmarshal the payload into a struct
// 			var data models.TSelektor
// 			err := json.Unmarshal(message.Body, &data)
// 			if err != nil {
// 				log.Printf("Error decoding message: %s\n", err)
// 				continue
// 			}

// 			log.Printf("Received message TSelektor:\n%+v\n", data)
// 			opts := options.Update().SetUpsert(true)
// 			filter := bson.D{{Key: "id", Value: data.IDSelektor}}
// 			update := bson.D{{Key: "$set", Value: data}}

// 			_, err = db.Database("mitra_tani_consumer").Collection("t_selektor_mtani").UpdateMany(context.TODO(), filter, update, opts)
// 			if err != nil {
// 				log.Printf("Error updating document: %s\n", err)
// 				continue
// 			}

// 		}

// 		defer channel.Close()
// 		defer db.Disconnect(context.TODO())
// 	}()

// 	go func() {
// 		for message := range msg_spta {
// 			// Unmarshal the payload into a struct
// 			var data models.TSpta
// 			err := json.Unmarshal(message.Body, &data)
// 			if err != nil {
// 				log.Printf("Error decoding message: %s\n", err)
// 				continue
// 			}

// 			log.Printf("Received message TSpta:\n%+v\n", data)
// 			opts := options.Update().SetUpsert(true)
// 			filter := bson.D{{Key: "id", Value: data.ID}}
// 			update := bson.D{{Key: "$set", Value: data}}

// 			_, err = db.Database("mitra_tani_consumer").Collection("t_spta_mtani").UpdateMany(context.TODO(), filter, update, opts)

// 			if err != nil {
// 				log.Printf("Error updating document: %s\n", err)
// 				continue
// 			}

// 		}

// 		defer channel.Close()
// 		defer db.Disconnect(context.TODO())
// 	}()

// 	go func() {
// 		for message := range msg_timbangan {
// 			// Unmarshal the payload into a struct
// 			var data models.TTimbang
// 			err := json.Unmarshal(message.Body, &data)
// 			if err != nil {
// 				log.Printf("Error decoding message: %s\n", err)
// 				continue
// 			}

// 			log.Printf("Received message TTimbang:\n%+v\n", data)
// 			opts := options.Update().SetUpsert(true)
// 			filter := bson.D{{Key: "id", Value: data.IDTimbangan}}
// 			update := bson.D{{Key: "$set", Value: data}}

// 			_, err = db.Database("mitra_tani_consumer").Collection("t_timbangan_mtani").UpdateMany(context.TODO(), filter, update, opts)

// 			if err != nil {
// 				log.Printf("Error updating document: %s\n", err)
// 				continue
// 			}

// 		}
// 		defer channel.Close()
// 		defer db.Disconnect(context.TODO())
// 	}()
// }

func MyFuncInsert(db *mongo.Client, channel *amqp.Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	msgs_ari, err := channel.Consume(
		"mitra_tani_ari", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msg_meja_tebu, err := channel.Consume(
		"mitra_tani_meja_tebu", // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msg_selektor, err := channel.Consume(
		"mitra_tani_selektor", // queue
		"",                    // consumer
		true,                  // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // arguments

	)

	if err != nil {
		log.Fatal(err)
	}

	msg_spta, err := channel.Consume(
		"mitra_tani_spta", // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // arguments
	)

	if err != nil {
		log.Fatal(err)
	}

	msg_timbangan, err := channel.Consume(
		"mitra_tani_timbangan", // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // arguments
	)

	if err != nil {
		log.Fatal(err)
	}
	SptaChan := make(chan models.TSpta)
	TSelectorChan := make(chan models.TSelektor)
	TTimbanganChan := make(chan models.TTimbang)
	MejaTebuChan := make(chan models.TMejaTebu)
	AriChan := make(chan models.TAri)
	go func() {
		for message := range msgs_ari {
			var data models.TAri
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Error decoding message: %s\n", err)
				continue
			}

			AriChan <- data
		}
	}()

	go func() {
		for message := range msg_meja_tebu {
			var data models.TMejaTebu
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Error decoding message: %s\n", err)
				continue
			}

			MejaTebuChan <- data
		}
	}()

	go func() {
		for message := range msg_selektor {
			var data models.TSelektor
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Error decoding message: %s\n", err)
				continue
			}

			TSelectorChan <- data
		}
	}()

	go func() {
		for message := range msg_spta {
			var data models.TSpta
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Error decoding message: %s\n", err)
				continue
			}

			SptaChan <- data
		}
	}()

	go func() {
		for message := range msg_timbangan {
			var data models.TTimbang
			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Error decoding message: %s\n", err)
				continue
			}

			TTimbanganChan <- data
		}
	}()

	for {
		select {
		case data := <-AriChan:
			log.Printf("Received message TAri :\n%+v\n", data)
			_, err := db.Database("mitra_tani_consumer").Collection("t_ari_mtani").InsertOne(context.Background(), data)
			if err != nil {
				log.Printf("Error updating document: %s\n", err)
			}

		case data := <-MejaTebuChan:
			log.Printf("Received message TMejaTebu:\n%+v\n", data)
			_, err := db.Database("mitra_tani_consumer").Collection("t_meja_tebu_mtani").InsertOne(context.Background(), data)
			if err != nil {
				log.Printf("Error updating document: %s\n", err)
			}

		case data := <-TSelectorChan:
			log.Printf("Received message TSelektor:\n%+v\n", data)
			_, err := db.Database("mitra_tani_consumer").Collection("t_selektor_mtani").InsertOne(context.Background(), data)
			if err != nil {
				log.Printf("Error updating document: %s\n", err)
			}

		case data := <-SptaChan:
			log.Printf("Received message TSpta:\n%+v\n", data)
			_, err := db.Database("mitra_tani_consumer").Collection("t_spta_mtani").InsertOne(context.Background(), data)
			if err != nil {
				log.Printf("Error updating document: %s\n", err)
			}

		case data := <-TTimbanganChan:
			log.Printf("Received message TTimbang:\n%+v\n", data)
			_, err := db.Database("mitra_tani_consumer").Collection("t_timbangan_mtani").InsertOne(context.Background(), data)
			if err != nil {
				log.Printf("Error updating document: %s\n", err)
			}
		}
	}
}
