package graphdb

import (
	"log"
	"testing"
)

//test set property on a node
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

//test update node properties
func TestUpdateNodeProperties(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := map[string]string{}
	data["name"] = "001"
	data["value01"] = "v01"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test UpdateNodeProperties")
	data2 := map[string]string{}
	data2["name"] = "002"
	err = session.UpdateNodeProperties(node.ID, data)
	if err != nil {
		t.Error(err)
	}
	log.Println("test UpdateNodeProperties finished")
	log.Println("clear data ...")
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
}

func TestGetPropertiesForNode(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := map[string]string{}
	data["name"] = "001"
	data["value01"] = "v01"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test GetPropertiesForNode")
	result, err := session.GetPropertiesForNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
	log.Println("test GetPropertiesForNode finished.")
	log.Println("clear data ...")
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
}

func TestDeletePropertyFromNode(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := map[string]string{}
	data["name"] = "001"
	data["value01"] = "v01"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test DeletePropertyFromNode")
	err = session.DeletePropertyFromNode(node.ID, "name")
	if err != nil {
		t.Error(err)
	}
	log.Println("test DeletePropertyFromNode finished.")
	log.Println("clear data ...")
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
}

func TestDeletePropertiesFromNode(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := map[string]string{}
	data["name"] = "001"
	data["value01"] = "v01"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test DeletePropertiesFromNode")
	err = session.DeletePropertiesFromNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("test DeletePropertiesFromNode finished.")
	log.Println("clear data ...")
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
}

func TestUpdateRelationshipProperties(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := map[string]string{}
	data["name"] = "001"
	data["value01"] = "v01"
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data2 := map[string]string{}
	data2["name"] = "002"
	data2["value01"] = "v02"
	node2, err := session.CreateNode(data2)
	if err != nil {
		t.Error(err)
	}
	relType := "LOVES"
	relDesc := map[string]string{}
	relDesc["Access"] = "+RW"
	relationship, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test UpdateRelationshipProperties")
	dataDesc := map[string]string{}
	dataDesc["Access"] = "+R"
	dataDesc["own"] = "david"
	err = session.UpdateRelationshipProperties(relationship[0].ID, dataDesc)
	if err != nil {
		t.Error(err)
	}
	log.Println("test UpdateRelationshipProperties finished.")
	log.Println("clear data ...")
	err = session.DeleteRelationship(relationship[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
}
