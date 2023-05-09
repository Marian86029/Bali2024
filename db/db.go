package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* 	MongoCN es el objeto de conexion a la DB */

var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://Marian86029:Dublin2023@cluster0.mc0aver.mongodb.net/")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	/*ConnectDB es la funcioon que me permite conectar la DB*/

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosoa a la DB")
	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
