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
	// err = session.CreateNodeIndex(indexName)
	indexType := "fulltext"
	indexProvider := "lucene"
	err = session.CreateNodeIndexWithConf(indexName, indexType, indexProvider)
	if err != nil {
		t.Error(err)
	}
	uniqueness := "create_or_fail"
	// oneNode := map[string]interface{}{
	// 	"name":     "Tobias",
	// 	"sequence": 1,
	// }
	// data := map[string]interface{}{
	// 	"key":        "name",
	// 	"value":      "Tobias",
	// 	"properties": oneNode,
	// }
	node1Data := map[string]interface{}{}
	node1Data["name"] = "Peter"
	node1Data["sequence"] = 2
	node1, err := session.CreateNode(node1Data)
	if err != nil {
		t.Error(err)
	}
	indexKey := "name"
	indexValue := "Peter"

	_, err = session.AddNodeToIndex(indexKey, indexValue, indexName, node1.ID)
	if err != nil {
		t.Error(err)
	}

	data := map[string]interface{}{}
	data["key"] = indexKey
	data["value"] = indexValue
	data["properties"] = node1Data
	result, err := session.CreateUniqueNode(indexName, uniqueness, data)
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
	// log.Println("clear data...")
	// err = session.DeleteNode(node1.ID)
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = session.DeleteNodeIndex(indexName)
	// if err != nil {
	// 	t.Error(err)
	// }
	log.Println("test CreateUniqueNode finshed!")
}
