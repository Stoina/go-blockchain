package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HandleRequest exported
// ...
func HandleRequest(writer http.ResponseWriter, request *http.Request, controller Controller) {

	log.Println("New request received...")
	log.Println("Request method: " + request.Method)

	requestMethod := ParseToRequestMethod(request.Method)

	if requestMethod != UNDEFINED {
		handleRequest(requestMethod, controller, writer, request)
	}
}

func handleRequest(method RequestMethod, rc Controller, writer http.ResponseWriter, request *http.Request) {
	if method == GET {
		handleGetRequest(rc, writer, request)
	} else if method == POST {
		handlePostRequest(rc, writer, request)
	}
}

func handleGetRequest(controller Controller, writer http.ResponseWriter, request *http.Request) {
	header := writer.Header()
	header.Add("content-type", "application/json; charset=UTF-8")
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	writer.Write([]byte(controller.Get(request.URL).ConvertToJSON()))
}

func handlePostRequest(controller Controller, writer http.ResponseWriter, request *http.Request) {

	content, err := getContentFromRequest(request)

	if err != nil {
		log.Fatal(err)
	}

	contentType := getContentTypeFromRequest(request)

	header := writer.Header()
	header.Add("content-type", "application/json")
	writer.Write([]byte(controller.Post(contentType, content).ConvertToJSON()))
}

func getContentFromRequest(request *http.Request) (string, error) {
	contentBytes, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}

func getContentTypeFromRequest(request *http.Request) string {
	contentType := request.Header.Get("Content-Type")
	return contentType
}
