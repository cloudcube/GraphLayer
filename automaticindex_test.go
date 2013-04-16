package graphdb

import (
	"log"
	"testing"
)

func TestFindNodeFromAutomaticIndexByMatch(t *testing.T) {
	log.Println("Starting test FindNodeFromAutomaticIndexByMatch")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"name": "I",
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	nodeKey := "name"
	nodeValue := "I"
	results, err := session.FindNodeFromAutomaticIndexByMatch(nodeKey, nodeValue)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(results))
	for _, result := range results {
		log.Println(result)
	}
	log.Println("clear data...")
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}

	//auto index can't be delete
	// key := "node_auto_index"
	// err = session.DeleteNodeIndex(key)
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestFindNodeFindNodeFromAutomaticIndexByQuery(t *testing.T) {
	log.Println("Starting test FindNodeFromAutomaticIndexByQuery")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"name": "I",
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "you"
	node2, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	nodeKey := "name"
	nodeValue := "*"
	results, err := session.FindNodeFromAutomaticIndexByQuery(nodeKey, nodeValue)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(results))
	for _, result := range results {
		log.Println(result)
	}
	log.Println("clear data...")
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
}
