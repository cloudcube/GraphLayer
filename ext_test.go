package graphdb

import (
	"log"
	"testing"
)

func TestGetServiceRoot(t *testing.T) {
	log.Println("Start testing GetServiceRoot")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println(session.URL)
	data, err := session.GetServiceRoot()
	extensions := data.Extensions
	log.Println("print extesion")
	log.Println(len(extensions))
	for _, extension := range extensions {
		log.Println(extension)
	}
	log.Println(data)
	log.Println("GetServiceRoot test finished!")
}
