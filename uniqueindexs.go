package graphdb

import (
	"encoding/json"
	"errors"
	"log"
)

// 创建唯一节点
func (session *Session) CreateUniqueNode(indexName string, uniqueness string, data map[string]interface{}) (result *GraphDataTemplate, err error) {
	session.Method = "post"
	url := session.URL + "/" + "index" + "/" + "node"
	if len(indexName) == 0 {
		return result, errors.New("IndexName is nil!")
	}
	url += "/" + indexName
	if len(uniqueness) == 0 {
		return result, errors.New("Uniquenes is nil")
	}
	url += "?" + "uniqueness=" + uniqueness
	log.Println(url)
	if data == nil {
		return result, errors.New("Node Infor is nil!")
	}
	log.Println(data)
	buf, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	log.Println(string(buf))
	body, err := session.Send(url, string(buf))
	if err != nil {
		return result, err
	}
	log.Println(session.StatusCode)
	results, err := session.Unmarshal(body)
	if err != nil {
		return result, err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	result = results[0]
	return result, session.NewError(errorList)
}
