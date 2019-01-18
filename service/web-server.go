package service

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

//StartWebServer ...
func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)
	logrus.Infof("Starting HTTP service at %v", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlers.LoggingHandler(os.Stdout, http.DefaultServeMux),
	}

	err := server.ListenAndServe()

	if err != nil {
		logrus.Panicf("error starting HTTP server at port %v: %v", port, err.Error())
	}
}
