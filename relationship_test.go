package GraphLayer

import (
	// "errors"
	"log"
	"testing"
)

func TestCreateRelationship(t *testing.T) {
	log.Println("Start creating Relationship!")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := make(map[string]interface{})
	data["name"] = "001"
	srcNode, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data = make(map[string]interface{})
	data["name"] = "002"
	destNode, err := session.CreateNode(data)

	relDesc := make(map[string]string)
	relDesc["name"] = "001"

	relType := "relationType"
	dataSet, err := session.CreateRelationship(srcNode.ID, destNode.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println(dataSet[0])
	log.Println("Create Relationship sucessful")
	log.Println("Clear data ...")
	err = session.DeleteRelationship(dataSet[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(srcNode.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(destNode.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("Data cleared")
	log.Println("Test CreateRelationship finished")
}

func TestGetRelationshipById(t *testing.T) {
	log.Println("Starting test GetRelationshipById")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare test data")
	data := make(map[string]interface{})
	data["name"] = "001"
	srcNode, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "002"
	destNode, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	relDesc := map[string]string{}
	relType := ""
	dataSet, err := session.CreateRelationship(srcNode.ID, destNode.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println("create rel result by id")
	log.Println(dataSet[0])
	log.Println("Starting test get relationship")
	dataSetResult, err := session.GetRelationshipById(dataSet[0].ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("search rel result by id")
	log.Println(dataSetResult[0])
	// if dataSet[0] != dataSetResult[0] {
	// 	t.Error(errors.New("关系查询错误！！！"))
	// }
	log.Println("Get relationship by id sucessful")
	log.Println("Clear data ...")
	err = session.DeleteRelationship(dataSetResult[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(srcNode.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(destNode.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("Data cleared")
	log.Println("Test GetRelationshipById finished")
}

func TestGetRelationshipsOnNode(t *testing.T) {
	log.Println("Starting test GetRelationshipsOnNode")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := make(map[string]interface{})
	data["name"] = "001"
	data1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "002"
	data2, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "003"
	data3, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "004"
	data4, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	relDesc := map[string]string{}
	relType := "LIKES"
	rel01, err := session.CreateRelationship(data1.ID, data2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relType = "LIKES"
	rel02, err := session.CreateRelationship(data2.ID, data3.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relType = "HATES"
	rel03, err := session.CreateRelationship(data2.ID, data4.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test get relationships")
	log.Println("get all relationships")
	typed := ""
	direction := "all"
	allRelationships, err := session.GetRelationshipsOnNode(data1.ID, typed, direction)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(allRelationships))
	log.Println("get incoming relationships")
	typed = ""
	direction = "in"
	incomingRel, err := session.GetRelationshipsOnNode(data2.ID, typed, direction)
	if err != nil {
		t.Error(err)
	}
	for _, incomingData := range incomingRel {
		log.Println(incomingData)
	}
	log.Println("get outgoing relationships")
	typed = ""
	direction = "out"
	outgoingRel, err := session.GetRelationshipsOnNode(data2.ID, typed, direction)
	if err != nil {
		t.Error(err)
	}
	for _, outgoingData := range outgoingRel {
		log.Println(outgoingData)
	}
	log.Println("get typed relationships")
	typed = "LIKES&HATES"
	direction = ""
	typedRel, err := session.GetRelationshipsOnNode(data2.ID, typed, direction)
	if err != nil {
		t.Error(err)
	}
	for _, typedData := range typedRel {
		log.Println(typedData)
	}
	log.Println("Clear Data ...")
	err = session.DeleteRelationship(rel01[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(rel02[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(rel03[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(data1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(data2.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(data3.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(data4.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("Data cleared")
	log.Println("test GetRelationshipsOnNode finished")
}

func TestGetRelationshipTypes(t *testing.T) {
	log.Println("Start test GetRelationshipTypes")
	log.Println("Prapare data ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := make(map[string]interface{})
	data["name"] = "001"
	node01, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "002"
	node02, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "003"
	node03, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	relDesc := make(map[string]string)
	relType := "LIKES"
	rel01, err := session.CreateRelationship(node01.ID, node02.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relType = "LOVES"
	rel02, err := session.CreateRelationship(node01.ID, node03.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relType = "HATES"
	rel03, err := session.CreateRelationship(node02.ID, node03.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println("Starting test get relationship types")
	typeResult, err := session.GetRelationshipTypes()
	if err != nil {
		t.Error(err)
	}
	for _, typeString := range typeResult {
		log.Println(typeString)
	}
	log.Println("Clear Data ...")
	err = session.DeleteRelationship(rel01[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(rel02[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(rel03[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node01.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node02.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node03.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("Data cleared")
	log.Println("test GetRelationshipTypes finished ")
}

func TestDeleteRelationship(t *testing.T) {
	log.Println("Start test DeleteRelationship")
	log.Println("Prapare data ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := make(map[string]interface{})
	data["name"] = "001"
	node01, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "002"
	node02, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	relDesc := make(map[string]string)
	relType := ""
	rel01, err := session.CreateRelationship(node01.ID, node02.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println("Clear data ...")
	err = session.DeleteRelationship(rel01[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node01.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node02.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
	log.Println("test DeleteRelationship finished")
}
