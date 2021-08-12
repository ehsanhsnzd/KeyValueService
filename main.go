package main

import (
	"biges/Repository"
	"biges/Services"
	"flag"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":9001", "http server address")

func main() {
	flag.Parse()

	//Inject repository to service
	dataRepository := Repository.NewData()
	api := &Services.SERVICE{
		DataRepository: dataRepository,
	}

	//Handle Routes
	http.HandleFunc("/set", api.HandleSet)
	http.HandleFunc("/get", api.HandleGet)


	//Store data in file at specific interval
	Services.Schedule(
		func() {
			Services.SaveData(dataRepository)
		}, 20 * time.Second)


	//Listen http port
	log.Fatal(http.ListenAndServe(*addr, nil))

}


