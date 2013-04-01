package graphdb

import (
	autoTest "github.com/remogatto/prettytest"
	"log"
	"testing"
)

//test setup start
type testSuite struct {
	autoTest.Suite
}

func TestRunner(t *testing.T) {
	autoTest.Run(t, new(testSuite))
}

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

func TestRemovePropertyFromRelationship(t *testing.T) {
	t.Log("Starting test RemovePropertyFromRelationship")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	t.Log("test data prapare.")
	data := map[string]string{}
	data["name"] = "node01"
	node1, err := session.CreateNode(data)
	data["name"] = "node02"
	node2, err := session.CreateNode(data)
	relDesc := map[string]string{}
	relDesc["Role"] = "Admin"
	relType := "LOVES"
	relationship, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	err = session.RemovePropertyFromRelationship(relationship[0].ID, "Role")
	if err != nil {
		t.Error(err)
	}
	t.Log("Starting Clear Data ...")
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
	t.Log("test Data Cleanred")
	t.Log("test RemovePropertyFromRelationship finished!")
}

func TestRemovePropertiesFromRelationship(t *testing.T) {
	t.Log("Starting test RemovePropertiesFromRelationship")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	t.Log("Test data prapare...")
	data := map[string]string{}
	data["name"] = "n01"

	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "n02"
	node2, err := session.CreateNode(data)
	if err != nil {

		t.Error(err)
	}
	relDesc := map[string]string{}
	relDesc["access"] = "+WR"
	relType := "ROLE"
	relationship, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	t.Log(relationship[0])
	err = session.RemovePropertiesFromRelationship(relationship[0].ID)
	if err != nil {
		t.Error(err)
	}
	t.Log("clear data...")
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
	t.Log("test data cleared")
	t.Log("test RemovePropertiesFromRelationship finished.")
}
