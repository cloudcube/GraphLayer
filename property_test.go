package graphdb

import (
	"errors"
	"log"
	"testing"
)

func TestGetPropertyByName(t *testing.T) {
	log.Println("Starting test GetPropertyByName")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := make(map[string]string)
	data["name"] = "001"
	node01, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	value, err := session.GetPropertyByName(node01.ID, "name")
	if err != nil {
		t.Error(err)
	}
	if value != "001" {
		errors.New("查询值错误")
	}
	log.Println("Test GetPropertyByName sucessful")
	log.Println("Clear data ...")
	err = session.DeleteNode(node01.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
	log.Println("test GetPropertyByName finished")
}

func TestSetPropertyOnNode(t *testing.T) {
	log.Println("Starting test SetPropertyOnNode")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := make(map[string]string)
	data["name"] = "001"
	node01, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "002"
	isReplace := false
	err = session.SetPropertyOnNode(node01.ID, data, isReplace)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "003"
	data["value01"] = "v01"
	data["value02"] = "v02"
	isReplace = true
	err = session.SetPropertyOnNode(node01.ID, data, isReplace)
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
