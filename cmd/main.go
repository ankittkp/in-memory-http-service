package main

import (
	"github.com/jinxankit/in-memory-http-service/internal"
	log "github.com/sirupsen/logrus"
)

func init() {
	//Set the log format to JSON format
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	log.Infof("Service Started")
	//config, err := configs.LoadConfig(".")
	//if err != nil {
	//	log.Fatalln("error loading config :", err)
	//}
	// build new application
	app, err := internal.NewApplication()
	if err != nil {
		log.Fatalln("error unable to create new application")
	}

	errs := make(chan error)
	go func() {
		errs <- app.StartHTTPServer()
	}()
	log.Infof("Listening on port 8080...")
	//TODO For internal Communications gprc server would be also good

	// If any of the HTTP or GRPC sever fails, exit
	log.Fatal(<-errs)
}
