package neo4j

import (
	// "errors"
	"fmt"
	// "reflect"
	"testing"
)

var (
	url      = "http://localhost:7474/db/data"
	user     = "david"
	password = "keepass"
)

func GetDbClient() (*Neo4j, error) {
	var (
		dbClient *Neo4j
		err      error
	)
	dbClient, err = NewNeo4j(url, user, password)
	if err != nil {
		fmt.Println(err)
	}
	return dbClient, err
}

/*
 * TestGetAllNodes test get all nodes
 */
/*func TestGetAllNodes(test *testing.T) {
	var nNode []string
	var err error
	var dbClient *Neo4j
	dbClient, err = NewNeo4j(url, user, password)
	if err != nil {
		test.Error(err.Error())
	}

	nNode, err = dbClient.GetAllNodes()
	if err != nil {
		test.Error(err.Error())
	}
	// fmt.Println(nNode)
	for _, node := range nNode {
		fmt.Println(node)
	}
}*/

/*type testNode struct {
	Name string
	Sex  string
	Age  int32
}

func struct2Map(data interface{}) (map[string]interface{}, error) {
	dataE := reflect.ValueOf(data)
	if dataE.Kind() == reflect.Ptr {

		dataE = dataE.Elem()
	}
	switch dataE.Kind() {
	case reflect.Struct:
		dataT := dataE.Type()
		result := make(map[string]interface{})
		for i := 0; i < dataE.NumField(); i++ {
			dataF := dataE.Field(i)
			result[dataT.Field(i).Name] = dataF.Interface()
		}

		return result, nil

	default:
		return nil, errors.New("不支持" + dataE.Type().Name() + "类型的转换")
	}
	return nil, nil
}

func TestNewNeo4j(t *testing.T) {
	_, err := NewNeo4j(url, user, password)
	if err != nil {
		t.Error(err)
	}
}

var testNodeData = []testNode{
	{"david2", "male", 30},
	{"xiaoshuomi2", "femal", 29},
	{"yangyang2", "male", 1},
}

func TestCreateNode(t *testing.T) {
	neo, err := NewNeo4j(url, user, password)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(testNodeData); i++ {
		result, errC := struct2Map(&testNodeData[i])
		if errC != nil {
			t.Error(err)
		}
		fmt.Println(result)
		neoData, errN := neo.CreateNode(result)
		if errN != nil {
			t.Error(errN)
		}
		fmt.Println(neoData.ID)
	}
}*/

func TestGetRelationTypes(test *testing.T) {
	var nRelationshipTypes []string
	dbClient, err := GetDbClient()
	nRelationshipTypes, err = dbClient.GetRelationTypes()
	if err != nil {
		test.Error(err.Error())
	}
	for _, relationshipType := range nRelationshipTypes {
		fmt.Println(relationshipType)
	}
}
