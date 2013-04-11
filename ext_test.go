package graphdb

import (
	"log"
	"strconv"
	"testing"
)

func TestGetServiceRoot(t *testing.T) {
	log.Println("Start testing GetServiceRoot")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println(session.URL)
	data, err := session.GetServiceRoot()
	extensions := data.Extensions
	log.Println("print extesion")
	log.Println(len(extensions))
	for _, extension := range extensions {
		log.Println(extension)
	}
	log.Println(data)
	log.Println("GetServiceRoot test finished!")
}

func TestCypherQueries(t *testing.T) {
	log.Println("Start testing CypherQueries")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare data ...")
	data := map[string]string{
		"name": "test01",
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "test02"
	node2, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "test03"
	data["age"] = "1"
	node3, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	relDesc := map[string]string{
		"v01": "v01",
	}
	relType := "KNOW"
	relationship1, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relDesc["v02"] = "v02"
	relationship2, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	log.Println("data ok!")
	query := "start x=node(" + strconv.FormatUint(node1.ID, 10) + ") match x-[r]->n return type(r),n.name?,n.age?"
	log.Println(query)
	parameters := map[string]string{}
	result, err := session.CypherQueries(query, parameters)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(result.Columns))
	for _, cloumn := range result.Columns {
		log.Println(cloumn)
	}
	for _, data := range result.Data {
		log.Println(data)
	}
	log.Println(relationship1[0].ID)
	log.Println(relationship2[0].ID)
	log.Println("clean data ...")
	err = session.DeleteRelationship(relationship1[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(relationship2[0].ID)
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
	err = session.DeleteNode(node3.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
	log.Println("CypherQueries test finished.")

}

func TestRemoveEntriesFromIndex(t *testing.T) {
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Prepare test data ...")
	log.Println("create index")
	indexName := "testIndex01"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("create one node")
	data := map[string]string{
		"name": "node01",
		"k01":  "v01",
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	indexKey := "indexKey01"
	indexValue := "indexValue01"
	_, err = session.AddNodeToIndex(indexKey, indexValue, indexName, node1.ID)
	if err != nil {
		t.Error(err)
	}
	data2 := map[string]string{
		"name": "node02",
		"k01":  "v01",
	}
	node2, err := session.CreateNode(data2)
	if err != nil {
		t.Error(err)
	}
	_, err = session.AddNodeToIndex(indexKey, indexValue, indexName, node2.ID)
	if err != nil {
		t.Error(err)
	}
	log.Println("start RemoveEntriesFromIndex given nodeId")
	key := ""
	value := ""
	err = session.RemoveEntriesFromIndex(node1.ID, indexName, key, value)
	if err != nil {
		t.Error(err)
	}
	results, err := session.FindNodeByMatch(indexName, indexKey, indexValue)
	log.Println(len(results))
	for _, result := range results {
		log.Println(result)
	}
	log.Println("test RemoveEntriesFromIndex given nodeId and indexKey")
	_, err = session.AddNodeToIndex(indexKey, indexValue, indexName, node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.RemoveEntriesFromIndex(node1.ID, indexName, indexKey, "")
	if err != nil {
		t.Error(err)
	}
	results, err = session.FindNodeByMatch(indexName, indexKey, indexValue)
	log.Println(len(results))
	for _, result := range results {
		log.Println(result)
	}
	log.Println("test RemoveEntriesFromIndex given nodeId,indexKey,indexValue")
	err = session.RemoveEntriesFromIndex(node2.ID, indexName, indexKey, indexValue)
	if err != nil {
		t.Error(err)
	}
	results, err = session.FindNodeByMatch(indexName, indexKey, indexValue)
	log.Println(len(results))
	log.Println("clean data ...")
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared")
	log.Println("RemoveEntriesFromIndex test finished!")

}

func TestFindNodeByMatch(t *testing.T) {
	log.Println("Starting test FindNodeByMatch")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	indexName := "testIndex01"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	data := map[string]string{}
	data["name"] = "testNode01"
	data["k01"] = "v01"
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	key := "indexKey01"
	value := "indexValue01"
	_, err = session.AddNodeToIndex(key, value, indexName, node1.ID)
	if err != nil {
		t.Error(err)
	}
	data1 := map[string]string{
		"name": "testNode02",
		"k01":  "v01",
		"k02":  "v02",
	}
	node2, err := session.CreateNode(data1)
	if err != nil {
		t.Error(err)
	}
	_, err = session.AddNodeToIndex(key, value, indexName, node2.ID)
	if err != nil {
		t.Error(err)
	}
	indexKey := key
	indexValue := value
	results, err := session.FindNodeByMatch(indexName, indexKey, indexValue)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(results))
	for _, result := range results {
		log.Println(result)
	}
	log.Println("clean data ...")
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleared!")
	log.Println("FindNodeByMatch test finished!")
}

func TestFindNodeByQuery(t *testing.T) {
	log.Println("Start testing FindNodeByQuery")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Create Index,type lucene")
	indexName := "testIndex"
	indexType := "fulltext"
	indexProvider := "lucene"
	err = session.CreateNodeIndexWithConf(indexName, indexType, indexProvider)
	if err != nil {
		t.Error(err)
	}
	log.Println("create two nodes")
	data := map[string]string{
		"name":  "test01",
		"key01": "value01",
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "test02"
	node2, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	indexKey := "Name"
	indexValue := "test name"
	_, err = session.AddNodeToIndex(indexKey, indexValue, indexName, node1.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = session.AddNodeToIndex(indexKey, indexValue, indexName, node2.ID)
	if err != nil {
		t.Error(err)
	}
	luceneQuery := "Name:test*"
	results, err := session.FindNodeByQuery(indexName, luceneQuery)
	// results, err := session.FindNodeByMatch(indexName, indexKey, indexValue)
	log.Println(len(results))
	for _, result := range results {
		log.Println(result)
	}
	log.Println("Clean data...")
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	log.Println("data cleaned")
	log.Println("FindNodeByQuery test finished!")
}
