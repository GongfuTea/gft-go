package mgo

import (
	"context"
	"log"

	"github.com/benweissmann/memongo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mgo struct {
	session *mongo.Client
}

var DefaultMongo = &Mgo{}

func OpenMongo() *Mgo {
	version := viper.GetString("mongo.memongo")
	if version != "" {
		return openMemongo()
	}

	host := viper.GetString("mongo.uri")
	db := viper.GetString("mongo.db")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    db,
		Username:      username,
		Password:      password,
	}
	clientOpts := options.Client().ApplyURI(host).SetAuth(credential)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	DefaultMongo.session = client
	return DefaultMongo
}

func openMemongo() *Mgo {
	version := viper.GetString("mongo.memongo")

	mongoServer, err := memongo.Start(version)
	if err != nil {
		log.Fatal(err)
	}
	// defer mongoServer.Stop()

	log.Println(mongoServer.URI())
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoServer.URI()))
	if err != nil {
		log.Fatal(err)
	}
	DefaultMongo.session = client
	return DefaultMongo
}

func (db Mgo) Close() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (db Mgo) Collection(collection string) *mongo.Collection {
	database := viper.GetString("mongo.db")

	return DefaultMongo.session.Database(database).Collection(collection)
}
