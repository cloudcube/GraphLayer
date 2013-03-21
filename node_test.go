package graphdb

import (
	"log"
	"testing"
)

var ID uint64

var settingFile = "databaseConfig.json"

func TestCreateNode(t *testing.T) {
	log.Println("Start testing CreateNode function")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := make(map[string]string)
	data["name"] = "001"
	data["sex"] = "male"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	ID = node.ID
	log.Println(node)
	log.Println("Create Node Sucessful!")
}

func TestGetNode(t *testing.T) {
	log.Println("Start testing GetNode function")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data, err := session.GetNode(ID)
	if err != nil {
		t.Error(err)
	}
	log.Println(data)
	log.Println("Get Node Sucessful")
}

func TestDeleteNode(t *testing.T) {
	log.Println("Start testing DeleteNode function")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("Delete Node Sucessful.")
}
