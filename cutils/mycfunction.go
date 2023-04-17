package main

import (
	"encoding/base64"
	"fmt"
	//"unsafe"
)

func MyFunction(name, password string) (string, error) {
	encodedPsk := base64.StdEncoding.EncodeToString([]byte(password))
	return encodedPsk, nil
}

func main() {
	fmt.Println("Plugin called")
}
