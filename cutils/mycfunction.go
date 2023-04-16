package main

/*
#include <stdlib.h>
#include <stdint.h>
size_t myfunction(const char *something, const uint8_t *text, size_t len, uint8_t *param2) {
	return 10;
}
*/
import "C"
import (
	"encoding/base64"
	"fmt"
	"unsafe"
)

func MyFunction(name, password string) (string, error) {
	mytext := make([]byte, 1024)

	cStr := C.CString("something" + C.GoString((*C.char)(unsafe.Pointer(&[]byte(name)[0]))))
	defer C.free(unsafe.Pointer(cStr))

	// define input and output parameters
	text := []byte(password)
	decodedLength := C.myfunction(cStr, (*C.uint8_t)(&text[0]), C.size_t(len(text)), (*C.uint8_t)(&mytext[0]))

	// convert C string to Go string and print output
	decodedPsk := string(mytext[:decodedLength])
	encodedPsk := base64.StdEncoding.EncodeToString([]byte(decodedPsk))
	return encodedPsk, nil
}

func main() {
	fmt.Println("Plugin called")
}
