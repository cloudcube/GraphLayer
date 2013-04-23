package graphdb

import (
	"log"
	"testing"
)

func TestTraversal(t *testing.T) {
	log.Println("Starting test Traversal!")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Parepare data ...")
	nodeIdSlice := make([]uint64, 10)
	data1 := map[string]interface{}{
		"name": "root",
	}
	node1, err := session.CreateNode(data1)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node1.ID)
	data2 := map[string]interface{}{
		"name": "johan",
	}
	node2, err := session.CreateNode(data2)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node2.ID)
	data3 := map[string]interface{}{
		"name": "Mattias",
	}
	node3, err := session.CreateNode(data3)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node3.ID)
	data4 := map[string]interface{}{
		"name": "Emil",
	}
	node4, err := session.CreateNode(data4)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node4.ID)
	data5 := map[string]interface{}{
		"name": "Peter",
	}
	node5, err := session.CreateNode(data5)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node5.ID)
	data6 := map[string]interface{}{
		"name": "Tobias",
	}
	node6, err := session.CreateNode(data6)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node6.ID)
	data7 := map[string]interface{}{
		"name": "Sara",
	}
	node7, err := session.CreateNode(data7)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice = append(nodeIdSlice, node7.ID)
	log.Println("create relationship!")
	relIdSlice := make([]uint64, 10)
	relDesc := map[string]string{}
	relType := "knows"
	rel1, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice = append(relIdSlice, rel1[0].ID)
	rel2, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice = append(relIdSlice, rel2[0].ID)
	rel3, err := session.CreateRelationship(node2.ID, node4.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice = append(relIdSlice, rel3[0].ID)
	rel4, err := session.CreateRelationship(node4.ID, node5.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice = append(relIdSlice, rel4[0].ID)
	rel5, err := session.CreateRelationship(node4.ID, node6.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice = append(relIdSlice, rel5[0].ID)
	relType = "loves"
	rel6, err := session.CreateRelationship(node6.ID, node7.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice = append(relIdSlice, rel6[0].ID)
	order := "breadth_first"
	return_filter := map[string]string{
		"body":     "position.endNode().getProperty('name').toLowerCase().contains('t')",
		"language": "javascript",
	}
	prune_evaluator := map[string]string{
		"body":     "position.length() > 10",
		"language": "javascript",
	}
	uniqueness := "node_global"
	relationships := []map[string]string{
		{
			"direction": "all",
			"type":      "knows",
		},
		{
			"direction": "all",
			"type":      "loves",
		},
	}
	var max_depth uint64
	max_depth = 3
	dataResults, err := session.Traversal(node1.ID, order, return_filter, prune_evaluator, uniqueness, relationships, max_depth)
	log.Println(len(dataResults))
	for _, result := range dataResults {
		log.Println(result)
	}
	log.Println("clear data...")
	log.Println("delete relationships")
	for _, relId := range relIdSlice {
		err = session.DeleteRelationship(relId)
		if err != nil {
			t.Error(err)
		}
	}
	log.Println("delelte nodes")
	for _, nodeId := range nodeIdSlice {
		err = session.DeleteNode(nodeId)
		if err != nil {
			t.Error(err)
		}
	}
	log.Println("data cleaned")
	log.Println("Traversal test finished!")
}
