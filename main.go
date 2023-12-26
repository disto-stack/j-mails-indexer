package main

import (
	"net/http"
	"os"

	"github.com/disto-stack/j-mails-indexer/pkg/handlers"
	"github.com/disto-stack/j-mails-indexer/pkg/services"
)

var (
	indexerHandler handlers.IndexerHandler
	
)
func main()  {
	setHandlersDependencies()

	args := os.Args
	if len(args) > 1 {
		filename := args[1]
		indexerHandler.IndexFromTgz(filename);
	}
}

func setHandlersDependencies() {
	config := &services.Config{}
	config.SetConfig("ssdsd")

	indexerHandler = handlers.IndexerHandler{}
	indexerHandler.SetDependencies(config)
}

func hello(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("Hello world"))
}