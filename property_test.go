package graphdb

import (
	"log"
	"testing"
)

func TestSetPropertyOnNode(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := make(map[string]string)
	data["name"] = "001"
	data["sex"] = "female"
	node01, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data02 := make(map[string]string)
	data02["name"] = "002"
	log.Println("Starting test SetPropertyOnNode")
	err = session.SetPropertyOnNode(node01.ID, data02)
	if err != nil {
		t.Error(err)
	}
	log.Println("Clear data ...")
	err = session.DeleteNode(node01.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
	log.Println("test SetPropertyOnNode finished.")
}
