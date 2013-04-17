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

func TestCrateAutoIndexWithConf(t *testing.T) {
	log.Println("Starting test CrateAutoIndexWithConf!")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("create auto index for node")
	indexType := "fulltext"
	indexProvider := "lucene"
	category := "node"
	err = session.CrateAutoIndexWithConf(indexType, indexProvider, category)
	if err != nil {
		t.Error(err)
	}
	log.Println("create auto index for relationship")
	category = "relationship"
	err = session.CrateAutoIndexWithConf(indexType, indexProvider, category)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAutoIndexStatus(t *testing.T) {
	log.Println("Starting test GetAutoIndexStatus")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("get node auto index status")
	category := "node"
	status, err := session.GetAutoIndexStatus(category)
	if err != nil {
		t.Error(err)
	}
	log.Println(status)
	log.Println("get relationship auto index status")
	category = "relationship"
	status, err = session.GetAutoIndexStatus(category)
	if err != nil {
		t.Error(err)
	}
	log.Println(status)
	log.Println("get auto index status test finished!")
}

func TestEnableAutoindex(t *testing.T) {
	log.Println("Starting test EnableNodeAutoindex")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("test set node auto index status")
	status := false
	category := "node"
	status, err = session.GetAutoIndexStatus(category)
	if err != nil {
		t.Error(err)
	}
	if !status {
		status = true
		err = session.EnableAutoindex(status, category)
		if err != nil {
			t.Error(err)
		}
	}
	status, err = session.GetAutoIndexStatus(category)
	if err != nil {
		t.Error(err)
	}
	log.Println("node auto idnex status")
	log.Println(status)
	log.Println("test set relationship auto index status")
	category = "relationship"
	status, err = session.GetAutoIndexStatus(category)
	if err != nil {
		t.Error(err)
	}
	if !status {
		log.Println("runing here!")
		status = true
		err = session.EnableAutoindex(status, category)
	}
	status, err = session.GetAutoIndexStatus(category)
	log.Println("relationship auto index status")
	log.Println(status)
	log.Println("EnableNodeAutoindex test finished")
}

func TestLookuplistAutoIndexProperties(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("test lookup list node autoindex properties")
	category := "node"
	results, err := session.LookuplistAutoIndexProperties(category)
	if err != nil {
		t.Error(err)
	}
	for _, result := range results {
		log.Println(result)
	}
	log.Println("loookup list relationship autoindex properties")
	category = "relationship"
	results, err = session.LookuplistAutoIndexProperties(category)
	if err != nil {
		t.Error(err)
	}
	for _, result := range results {
		log.Println(result)
	}
	log.Println("LookuplistAutoIndexProperties test finished!")
}

func TestAddPropertyForAutoIndex(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("test add property for node autoindex ")
	category := "node"
	property := "education"
	err = session.AddPropertyForAutoIndex(property, category)
	if err != nil {
		t.Error(err)
	}
	log.Println("lookup list node autoindex properties")
	results, err := session.LookuplistAutoIndexProperties(category)
	for _, result := range results {
		log.Println(result)
	}
	log.Println("test add property for relationship autoindex")
	category = "relationship"
	property = "life"
	err = session.AddPropertyForAutoIndex(property, category)
	if err != nil {
		t.Error(err)
	}
	log.Println("lookup list relationship autoindex properties")
	results, err = session.LookuplistAutoIndexProperties(category)
	if err != nil {
		t.Error(err)
	}
	for _, result := range results {
		log.Println(result)
	}
	log.Println("AddPropertyForAutoIndex test finished!")
}

func TestRemovePropertyForAutoIndex(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("test remove property for node autoindex")
	category := "node"
	property := "book"
	err = session.RemovePropertyForAutoIndex(property, category)
	if err != nil {
		t.Error(err)
	}
	log.Println("lookup list node autoindex properties")
	results, err := session.LookuplistAutoIndexProperties(category)
	for _, result := range results {
		log.Println(result)
	}
	log.Println("test remove property for relationship autoindex")
	property = "life"
	category = "relationship"
	err = session.RemovePropertyForAutoIndex(property, category)
	if err != nil {
		t.Error(err)
	}
	log.Println("lookup list relationship autoindex properties")
	results, err = session.LookuplistAutoIndexProperties(category)
	if err != nil {
		t.Error(err)
	}
	for _, result := range results {
		log.Println(result)
	}
	log.Println("RemovePropertyForAutoIndex test finished!")

}
