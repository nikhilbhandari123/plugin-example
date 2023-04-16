package main

import (
	"github.com/nikhilbhandari123/plugin-example/internal/routers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":65008", router))
}
