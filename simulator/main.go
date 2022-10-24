package main

import (
	"log"
	"simulator/infra/kafka"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProcucer()
	kafka.Publish("ola", "readtest", producer)
	for {
		_ = 1
	}

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()

	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[1])
}
