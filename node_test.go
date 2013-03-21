package graphdb

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	session, err := Dial("databaseConfig.json")
	if err != nil {
		t.Error(err)
	}
	data := make(map[string]string)
	data["name"] = "001"
	data["sex"] = "male"
	node, err := session.CreateNode(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(node)
}
