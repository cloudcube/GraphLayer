package graphdb

import (
	"log"
	"testing"
)

func TestCreateNodeIndex(t *testing.T) {
	log.Println("test createnodeindex ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]string{}
	data["name"] = "index01"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	// log.Println(node)
	//data's key must be name
	// data["name"] = "test"
	indexName := "indexName001"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("Clear data ...")
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared!")
	log.Println("CreateNodeIndex finished!")
}

func TestCreateNodeIndexWithConf(t *testing.T) {
	log.Println("test createnodeindexwith config ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]string{}
	data["name"] = "index01"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	// log.Println(node)
	//data's key must be name
	// data["name"] = "test"
	indexName := "indexName001"
	indexType := "exact"
	indexProvider := "lucene"
	err = session.CreateNodeIndexWithConf(indexName, indexType, indexProvider)
	if err != nil {
		t.Error(err)
	}
	log.Println("Clear data ...")
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared!")
	log.Println("CreateNodeIndex finished!")
}

func TestListNodeIndexes(t *testing.T) {
	log.Println("test ListNodeIndexes ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	// data := map[string]string{}
	// data["name"] = "test001"
	indexName := "indexName001"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	indexResult, err := session.ListNodeIndexes()
	if err != nil {
		t.Error(err)
	}
	log.Println(indexResult)
	log.Println("clear data ...")
	// key := "test001"
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
	log.Println("ListNodeIndexes finished")
}

func TestDeleteNodeIndex(t *testing.T) {
	log.Println("test DeleteNodeIndex ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	// data := map[string]string{}
	// data["name"] = "test002"
	indexName := "indexName001"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	// key := "test002"
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("DeleteNodeIndex finished")
}

func TestAddNodeToIndex(t *testing.T) {
	log.Println("test AddNodeToIndex ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]string{}
	data["name"] = "v01"
	node01, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	// indexData := map[string]string{}
	// indexData["name"] = "testIndex"
	indexName := "indexName001"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "v02"
	node02, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	key := "some key"
	value := "some value"
	result, err := session.AddNodeToIndex(key, value, indexName, node02.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("index result:...")
	log.Println(result)
	indexResult, err := session.ListNodeIndexes()
	if err != nil {
		t.Error(err)
	}
	log.Println(indexResult)
	log.Println("clear data ...")
	err = session.DeleteNodeIndex(indexName)
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
	log.Println("AddNodeToIndex test finished.")
}

func TestCreateAutoIndexForNodes(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]string{
		"name": "autoIndex",
	}
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	indexName := "node_auto_index"
	indexType := "fulltext"
	indexProvider := "lucene"
	err = session.CreateAutoIndexForNodes(indexName, indexType, indexProvider)
	if err != nil {
		t.Error(err)
	}
	log.Println("clear data ...")
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared!")
}

func TestCreateAutoIndexForRelationships(t *testing.T) {
	log.Println("start testing CreateAutoIndexForRelationships ...")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	indexName := "relationship_auto_index"
	indexType := "exact"
	indexProvider := "lucene"
	err = session.CreateAutoIndexForRelationships(indexName, indexType, indexProvider)
	if err != nil {
		t.Error(err)
	}
	log.Println("clear data ...")
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("dat cleared!")
	log.Println("CreateAutoIndexForRelationships test finished!")
}
