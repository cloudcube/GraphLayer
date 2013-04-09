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
	data := map[string]string{
		"name": "kn01",
		"v01":  "v01",
		"v02":  "v02",
	}
	node1, err := session.CreateNode(data)
	indexName := "kvnode"
	err = session.CreateNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	key := ""
	value := ""
	err = session.RemoveEntriesFromIndex(node1.ID, indexName, key, value)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node1.ID)
	if err != nil {
		t.Error(err)
	}

	data2 := map[string]string{
		"name": "kn02",
		"v01":  "v01",
		"v02":  "v02",
	}
	node2, err := session.CreateNode(data2)
	key = "v01"
	value = ""
	err = session.RemoveEntriesFromIndex(node2.ID, indexName, key, value)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node2.ID)
	if err != nil {
		t.Error(err)
	}
	node3, err := session.CreateNode(data2)
	key = "v01"
	value = "v01"
	err = session.RemoveEntriesFromIndex(node3.ID, indexName, key, value)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNodeIndex(indexName)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteNode(node3.ID)
	if err != nil {
		t.Error(err)
	}
}
