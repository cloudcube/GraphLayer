package graphdb

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
	result, err := session.CreateUniqueNode(indexName, data)
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
	log.Println("Create a unique node in an index (the case where it exists)")
	data1 := map[string]interface{}{
		"key":        "okokok",
		"value":      "Exist Node Test.",
		"properties": oneNode,
	}
	result, err = session.CreateUniqueNode(indexName, data1)
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
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
	result, err = session.CreateUniqueNode(indexName, data2)
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
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
