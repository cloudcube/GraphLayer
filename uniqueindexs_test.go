package goneo4j

import (
	"log"
	"testing"
)

func TestCreateUniqueNode(t *testing.T) {
	log.Println("Start testing Get or create unique node (create)")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	indexName := "people"
	indexType := "fulltext"
	indexProvider := "lucene"
	err = session.CreateNodeIndexWithConf(indexName, indexType, indexProvider)
	if err != nil {
		t.Error(err)
	}
	// uniqueness := "unique"
	oneNode := map[string]interface{}{
		"name":     "Tobias",
		"sequence": 2,
	}
	data := map[string]interface{}{
		"key":        "name",
		"value":      "Tobias",
		"properties": oneNode,
	}
	log.Println("Create a unique node in an index")
	result1, err := session.CreateUniqueNode(indexName, data)
	if err != nil {
		t.Error(err)
	}
	log.Println("&&&&&&&&&&&&&&&")
	log.Println(result1.ID)
	log.Println("Create a unique node in an index (the case where it exists)")
	data1 := map[string]interface{}{
		"key":        "name",
		"value":      "Tobias",
		"properties": oneNode,
	}
	result2, err := session.CreateUniqueNode(indexName, data1)
	if err != nil {
		t.Error(err)
	}
	log.Println(result2)
	log.Println("Add a node to an index unless a node already exists for the given mapping")
	nodeData := map[string]interface{}{
		"name": "node002",
		"key":  100,
	}
	node02, err := session.CreateNode(nodeData)
	if err != nil {
		t.Error(err)
	}
	data2 := map[string]interface{}{
		"key":   "name",
		"value": "Mattias",
		"uri":   node02.Self,
	}
	result3, err := session.CreateUniqueNode(indexName, data2)
	if err != nil {
		t.Error(err)
	}
	log.Println(result3)
	log.Println("clear data ...")
	err = session.DeleteNode(node02.ID)
	if err != nil {
		t.Error(err)
	}
	indexKey := "name"
	indexValue := "Tobias"
	results, err := session.FindNodeByMatch(indexName, indexKey, indexValue)
	if err != nil {
		t.Error(err)
	}
	for _, dataResult := range results {
		log.Println(dataResult)
		log.Println("node Id:")
		log.Println(dataResult.ID)
		err = session.DeleteNode(dataResult.ID)
		if err != nil {
			t.Error(err)
		}
	}
	key := indexKey
	err = session.DeleteNodeIndex(key)
	if err != nil {
		t.Error(err)
	}
	key = "okokok"
	err = session.DeleteNodeIndex(key)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleaned!")
	log.Println("test CreateUniqueNode finshed!")
}

func TestCreateUniqueRelationship(t *testing.T) {
	log.Println("Create a unique relationship in an index")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("prepare data...")
	data1 := map[string]interface{}{
		"name": "node01",
	}
	node1, err := session.CreateNode(data1)
	if err != nil {
		t.Error(err)
	}
	data2 := map[string]interface{}{
		"name": "node02",
	}
	node2, err := session.CreateNode(data2)
	if err != nil {
		t.Error(err)
	}
	indexKey := "index01"
	indexValue := "value01"
	startUri := node1.Self
	endUri := node2.Self
	relationType := "knowledge"
	resultData, err := session.CreateUniqueRelationship(indexKey, indexValue, startUri, endUri, relationType)
	if err != nil {
		t.Error(err)
	}
	log.Println(resultData)
	log.Println("clear data...")
	err = session.DeleteRelationship(resultData.ID)
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
	err = session.DeleteRelationshipIndex(relationType)
	if err != nil {
		t.Error(err)
	}
	//but i don't know how to delete relationship's index.
	log.Println("data cleaned!")
}

func TestAddRelationshipToIndex(t *testing.T) {
	log.Println("Add a relationship to an index unless a relationship already exists for the given mapping")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("prepare data...")
	data1 := map[string]interface{}{
		"name": "node01",
	}
	node1, err := session.CreateNode(data1)
	if err != nil {
		t.Error(err)
	}
	data2 := map[string]interface{}{
		"name": "node02",
	}
	node2, err := session.CreateNode(data2)
	if err != nil {
		t.Error(err)
	}
	relDesc := map[string]string{
		"name": "rel01",
		"k01":  "v01",
	}
	relType := "DENY"
	dataSet, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)

	indexKey := "k01"
	indexValue := "v01"
	result, err := session.AddRelationshipToIndex(indexKey, indexValue, dataSet[0].Self, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
	// log.Println(resultData)
	log.Println("clear data...")
	err = session.DeleteRelationship(result.ID)
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
	err = session.DeleteRelationshipIndex(relType)
	if err != nil {
		t.Error(err)
	}
	//but i don't know how to delete relationship's index.
	log.Println("data cleaned!")
	log.Println("test finished!!!")
}
