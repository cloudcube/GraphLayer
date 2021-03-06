package goneo4j

import (
	"log"
	// "math/rand"
	"testing"
	// "time"
	// "strings"
	"strconv"
)

func TestTraversal(t *testing.T) {
	log.Println("Starting test Traversal!")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	log.Println("Parepare data ...")
	nodeIdSlice := make([]uint64, 7)
	data1 := map[string]interface{}{
		"name": "root",
	}
	node1, err := session.CreateNode(data1)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[0] = node1.ID
	data2 := map[string]interface{}{
		"name": "johan",
	}
	node2, err := session.CreateNode(data2)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[1] = node2.ID
	data3 := map[string]interface{}{
		"name": "Mattias",
	}
	node3, err := session.CreateNode(data3)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[2] = node3.ID
	data4 := map[string]interface{}{
		"name": "Emil",
	}
	node4, err := session.CreateNode(data4)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[3] = node4.ID
	data5 := map[string]interface{}{
		"name": "Peter",
	}
	node5, err := session.CreateNode(data5)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[4] = node5.ID
	data6 := map[string]interface{}{
		"name": "Tobias",
	}
	node6, err := session.CreateNode(data6)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[5] = node6.ID
	data7 := map[string]interface{}{
		"name": "Sara",
	}
	node7, err := session.CreateNode(data7)
	if err != nil {
		t.Error(err)
	}
	nodeIdSlice[6] = node7.ID
	log.Println("create relationship!")
	relIdSlice := make([]uint64, 6)
	relDesc := map[string]string{}
	relType := "knows"
	rel1, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice[0] = rel1[0].ID
	rel2, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice[1] = rel2[0].ID
	rel3, err := session.CreateRelationship(node2.ID, node4.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice[2] = rel3[0].ID
	rel4, err := session.CreateRelationship(node4.ID, node5.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice[3] = rel4[0].ID
	rel5, err := session.CreateRelationship(node4.ID, node6.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice[4] = rel5[0].ID
	relType = "loves"
	rel6, err := session.CreateRelationship(node6.ID, node7.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relIdSlice[5] = rel6[0].ID
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
	dataResults, err := session.TraversalByFilter(node1.ID, order, return_filter, prune_evaluator, uniqueness, relationships, max_depth)
	log.Println(len(dataResults))
	for _, result := range dataResults {
		log.Println(result)
	}
	log.Println("clear data...")
	log.Println("delete relationships")
	log.Println(len(relIdSlice))
	for _, relId := range relIdSlice {
		log.Println(relId)
		err = session.DeleteRelationship(relId)
		if err != nil {
			t.Error(err)
		}
	}
	log.Println("delelte nodes")
	log.Println(len(nodeIdSlice))
	for _, nodeId := range nodeIdSlice {
		log.Println(nodeId)
		err = session.DeleteNode(nodeId)
		if err != nil {
			t.Error(err)
		}
	}
	log.Println("data cleaned")
	log.Println("Traversal test finished!")
}

func TestGetRelationshipsFromTraversal(t *testing.T) {
	log.Println("Start testing return relationships from a traversal!")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"name": 'I',
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "car"
	node2, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "you"
	node3, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	log.Println("Create relationships")
	relDesc := map[string]string{}
	relType := "know"
	rel1, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relType = "own"
	rel2, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	order := "breadth_first"
	uniqueness := "none"
	return_filter := map[string]string{
		"language": "builtin",
		"name":     "all",
	}
	dataResults, err := session.GetRelationshipsFromTraversal(node1.ID, order, uniqueness, return_filter)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(dataResults))
	for _, dataResult := range dataResults {
		log.Println(dataResult)
	}
	log.Println("delete relationships")
	err = session.DeleteRelationship(rel1[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(rel2[0].ID)
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
	log.Println("data cleanred!")
}

func TestGetPathsFromTraversal(t *testing.T) {
	log.Println("Start testing return relationships from a traversal!")
	session, err := Dial(settingFile)
	if err != nil {
		t.Error(err)
	}
	data := map[string]interface{}{
		"name": 'I',
	}
	node1, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "car"
	node2, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	data["name"] = "you"
	node3, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	log.Println("Create relationships")
	relDesc := map[string]string{}
	relType := "know"
	rel1, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	relType = "own"
	rel2, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	if err != nil {
		t.Error(err)
	}
	order := "breadth_first"
	uniqueness := "none"
	return_filter := map[string]string{
		"language": "builtin",
		"name":     "all",
	}
	dataResults, err := session.GetPathsFromTraversal(node1.ID, order, uniqueness, return_filter)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(dataResults))
	for _, dataResult := range dataResults {
		log.Println(dataResult)
	}
	log.Println("delete relationships")
	err = session.DeleteRelationship(rel1[0].ID)
	if err != nil {
		t.Error(err)
	}
	err = session.DeleteRelationship(rel2[0].ID)
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
	log.Println("data cleanred!")
}

func TestGetNodesBelowDepthAtTraversal(t *testing.T) {
	session, err := Dial(settingFile)
	checkError(err, t)
	log.Println("Prepare data...")
	data := map[string]interface{}{}
	data["name"] = "Root"
	node1, err := session.CreateNode(data)
	checkError(err, t)
	data["name"] = "Johan"
	node2, err := session.CreateNode(data)
	checkError(err, t)
	relDesc := map[string]string{}
	relType := "knows"
	rel1, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	checkError(err, t)
	data["name"] = "Mattias"
	node3, err := session.CreateNode(data)
	checkError(err, t)
	rel2, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	checkError(err, t)
	data["name"] = "Emil"
	node4, err := session.CreateNode(data)
	checkError(err, t)
	rel3, err := session.CreateRelationship(node2.ID, node4.ID, relDesc, relType)
	checkError(err, t)
	data["name"] = "Peter"
	node5, err := session.CreateNode(data)
	checkError(err, t)
	rel4, err := session.CreateRelationship(node4.ID, node5.ID, relDesc, relType)
	checkError(err, t)
	data["name"] = "Tobias"
	node6, err := session.CreateNode(data)
	checkError(err, t)
	rel5, err := session.CreateRelationship(node4.ID, node6.ID, relDesc, relType)
	checkError(err, t)
	data["name"] = "Sara"
	relType = "loves"
	node7, err := session.CreateNode(data)
	checkError(err, t)
	rel6, err := session.CreateRelationship(node6.ID, node7.ID, relDesc, relType)
	checkError(err, t)
	log.Println("data already!!!")
	log.Println("Traversal returning nodes below a certain depth")
	return_filter := map[string]string{
		"body":     "position.length()<3;",
		"language": "javascript",
	}
	prune_evaluator := map[string]string{
		"name":     "none",
		"language": "builtin",
	}
	dataResults, err := session.GetNodesBelowDepthAtTraversal(node1.ID, return_filter, prune_evaluator)
	checkError(err, t)
	log.Println(len(dataResults))
	for key, dataResult := range dataResults {
		log.Println(key)
		log.Println(dataResult)
	}
	log.Println("Clean data...")
	err = session.DeleteRelationship(rel1[0].ID)
	checkError(err, t)
	err = session.DeleteRelationship(rel2[0].ID)
	checkError(err, t)
	err = session.DeleteRelationship(rel3[0].ID)
	checkError(err, t)
	err = session.DeleteRelationship(rel4[0].ID)
	checkError(err, t)
	err = session.DeleteRelationship(rel5[0].ID)
	checkError(err, t)
	err = session.DeleteRelationship(rel6[0].ID)
	checkError(err, t)
	err = session.DeleteNode(node1.ID)
	checkError(err, t)
	err = session.DeleteNode(node2.ID)
	checkError(err, t)
	err = session.DeleteNode(node3.ID)
	checkError(err, t)
	err = session.DeleteNode(node4.ID)
	checkError(err, t)
	err = session.DeleteNode(node5.ID)
	checkError(err, t)
	err = session.DeleteNode(node6.ID)
	checkError(err, t)
	err = session.DeleteNode(node7.ID)
	checkError(err, t)
	log.Println("data cleaned!!!")
	log.Println("test finished!!!")
}

func TestCreatedPagedTraversers(t *testing.T) {
	session, err := Dial(settingFile)
	checkError(err, t)
	nodeIds, relIds := session.createData(t)
	log.Println("start testing CreatedPagedTraversers")
	prune_evaluator := map[string]string{
		"language": "builtin",
		"name":     "none",
	}
	return_filter := map[string]string{
		"language": "javascript",
		"body":     "position.endNode().getProperty('name').contains('1');",
	}
	order := "depth_first"
	relationships := map[string]string{
		"type":      "NEXT",
		"direction": "out",
	}
	dataResults, err := session.CreatedPagedTraversers(nodeIds[0], prune_evaluator, return_filter, order, relationships)
	checkError(err, t)
	log.Println(len(dataResults))
	for key, dataResult := range dataResults {
		log.Println(key)
		log.Println(dataResult)
	}
	log.Println(session.Location)
	log.Println("Clean data...")
	err = session.cleanData(nodeIds, relIds)
	checkError(err, t)
	log.Println("data cleaned!!!")
	log.Println("test finished!!!")
}

func TestGetPagingResultPagedTraverser(t *testing.T) {
	session, err := Dial(settingFile)
	checkError(err, t)
	log.Println("create nodeIds")
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nodeIds := []uint64{}
	var index int64 = 1
	var nodeNames [31]int64
	for i := 0; i < 31; i++ {
		// nodeNames[i] = r.Int63n(1000)
		nodeNames[i] = index
		index++
	}
	for _, nodeName := range nodeNames {
		data := map[string]interface{}{}
		data["name"] = strconv.FormatInt(nodeName, 10)
		node, err := session.CreateNode(data)
		if err != nil {
			t.Error(err)
		}
		nodeIds = append(nodeIds, node.ID)
	}
	log.Println("nodeIds created")
	log.Println("start testing CreatedPagedTraversers")
	prune_evaluator := map[string]string{
		"language": "builtin",
		"name":     "none",
	}
	return_filter := map[string]string{
		"language": "javascript",
		"body":     "position.endNode().getProperty('name').contains('1');",
	}
	order := "depth_first"
	relationships := map[string]string{
		"type":      "NEXT",
		"direction": "out",
	}
	log.Println(nodeIds[0])
	_, err = session.CreatedPagedTraversers(nodeIds[0], prune_evaluator, return_filter, order, relationships)
	if err != nil {
		t.Error(err)
	}
	log.Println(session.Location)
	// log.Println("starting test GetPagingResultPagedTraverser")
	// dataSets, err := session.GetPagingResultPagedTraverser(session.Location)
	// if err != nil {
	// 	t.Error(err)
	// }
	// for _, dataSet := range dataSets {
	// 	log.Println(dataSet)
	// }
	// log.Println("clean data...")
	// err = session.cleanData(nodeIds, relIds)
	// checkError(err, t)
	// log.Println("data cleaned")
	// log.Println("test finished!")
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

func (session *Session) createData(t *testing.T) ([]uint64, []uint64) {
	log.Println("Prepare data...")
	nodeIds := []uint64{}
	relIds := []uint64{}
	data := map[string]interface{}{}
	data["name"] = "13"
	node1, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node1.ID)
	data["name"] = "19"
	node2, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node2.ID)
	relDesc := map[string]string{}
	relType := "knows"
	rel1, err := session.CreateRelationship(node1.ID, node2.ID, relDesc, relType)
	checkError(err, t)
	relIds = append(relIds, rel1[0].ID)
	data["name"] = "17"
	node3, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node3.ID)
	rel2, err := session.CreateRelationship(node1.ID, node3.ID, relDesc, relType)
	checkError(err, t)
	relIds = append(relIds, rel2[0].ID)
	data["name"] = "31"
	node4, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node4.ID)
	rel3, err := session.CreateRelationship(node2.ID, node4.ID, relDesc, relType)
	checkError(err, t)
	relIds = append(relIds, rel3[0].ID)
	data["name"] = "21"
	node5, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node5.ID)
	rel4, err := session.CreateRelationship(node4.ID, node5.ID, relDesc, relType)
	checkError(err, t)
	relIds = append(relIds, rel4[0].ID)
	data["name"] = "55"
	node6, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node6.ID)
	rel5, err := session.CreateRelationship(node4.ID, node6.ID, relDesc, relType)
	checkError(err, t)
	relIds = append(relIds, rel5[0].ID)
	data["name"] = "99"
	relType = "loves"
	node7, err := session.CreateNode(data)
	checkError(err, t)
	nodeIds = append(nodeIds, node7.ID)
	rel6, err := session.CreateRelationship(node6.ID, node7.ID, relDesc, relType)
	checkError(err, t)
	relIds = append(relIds, rel6[0].ID)
	return nodeIds, relIds
}

func (session *Session) cleanData(nodeIds []uint64, relIds []uint64) (err error) {
	for _, relId := range relIds {
		err = session.DeleteRelationship(relId)
		if err != nil {
			return err
		}
	}
	for _, nodeId := range nodeIds {
		err = session.DeleteNode(nodeId)
		if err != nil {
			return err
		}
	}
	return nil
}
