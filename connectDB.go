package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* 	MongoCN es el objeto de conexion a la DB */

var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://marianoromera186:Bali2024>@myfirstproject.w1hlyqg.mongodb.net/?retryWrites=true&w=majority")

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
