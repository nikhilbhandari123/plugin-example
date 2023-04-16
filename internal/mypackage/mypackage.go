package mypackage

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"plugin"
)

type PasswordRequest struct {
	Password string `json:"certenckey"`
}

type TemplateData struct {
	Psk      string
	UserName string
	CertData string
}

var (
	PluginPath = "/lib/myplugin.so"
)

func getMyFunction(name, pwd string) (string, error) {
	// load the plugin
	plug, err := plugin.Open(PluginPath)
	if err != nil {
		log.Errorf("Error in opening plugin: %v", err)
		return "", err
		// handle error
	}

	// get a handle to the GetPskEncode function
	sym, err := plug.Lookup("MyFunction")
	if err != nil {
		log.Errorf("Error in finding function in plugin: %v", err)
		return "", err
		// handle error
	}
	// convert the symbol to the appropriate function type
	myFunction, ok := sym.(func(string, string) (string, error))
	if !ok {
		log.Errorf("Error in finding symbol in plugin: %v", err)
		return "", errors.New("Invalid symbol in plugin")
	}
	return myFunction(name, pwd)
}

func HttpCaller(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a PasswordRequest struct
	mydata, err := getMyFunction("name", "test")
	if err != nil {
		http.Error(w, "Cannot decode PSK", http.StatusInternalServerError)
		return
	}
	log.Infof("Got data back %v", mydata)

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename=iOS_IPSECProfile.mobileconfig")
	w.Header().Set("Content-Type", "application/octet-stream")
}
