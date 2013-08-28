package goneo4j

import (
	"encoding/json"
	"errors"
	// "log"
)

// 创建唯一节点
func (session *Session) CreateUniqueNode(indexName string /*uniqueness string,*/, data map[string]interface{}) (result *GraphDataTemplate, err error) {
	session.Method = "post"
	url := session.URL + "/" + "index" + "/" + "node"
	if len(indexName) == 0 {
		return result, errors.New("IndexName is nil!")
	}
	url += "/" + indexName
	// if len(uniqueness) == 0 {
	// 	return result, errors.New("Uniquenes is nil")
	// }
	//url += "?" + "uniqueness=" + uniqueness
	url += "?unique"
	// log.Println(url)
	if data == nil {
		return result, errors.New("Node Infor is nil!")
	}
	// log.Println(data)
	buf, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	// log.Println(string(buf))
	body, err := session.Send(url, string(buf))
	if err != nil {
		return result, err
	}
	// log.Println(session.StatusCode)
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

func (session *Session) CreateUniqueRelationship(indexKey string, indexValue string, startUri string, endUri string, relationType string) (result *GraphDataTemplate, err error) {
	session.Method = "post"
	url := session.URL + "/" + "index" + "/" + "relationship"
	if len(relationType) == 0 {
		return result, errors.New("relationType is nil!")
	}
	url += "/" + relationType
	url += "/" + "?unique"
	// log.Println(url)
	if len(indexKey) == 0 || len(indexValue) == 0 || len(startUri) == 0 || len(endUri) == 0 {
		return result, errors.New("Invalid data ...")
	}
	data := map[string]string{
		"key":   indexKey,
		"value": indexValue,
		"start": startUri,
		"end":   endUri,
		"type":  relationType,
	}
	// log.Println(data)
	buf, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	// log.Println(string(buf))
	body, err := session.Send(url, string(buf))
	// log.Println(body)
	if err != nil {
		return result, err
	}
	// log.Println(session.StatusCode)
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

func (session *Session) AddRelationshipToIndex(indexKey string, indexValue string, relationshipUri string, relationType string) (result *GraphDataTemplate, err error) {
	session.Method = "post"
	url := session.URL + "/" + "index" + "/" + "relationship"
	if len(relationType) == 0 {
		return result, errors.New("relationType is nil!")
	}
	url += "/" + relationType
	url += "/" + "?unique"
	// log.Println(url)
	if len(indexKey) == 0 || len(indexValue) == 0 || len(relationshipUri) == 0 {
		return result, errors.New("Invalid data ...")
	}
	data := map[string]string{
		"key":   indexKey,
		"value": indexValue,
		"uri":   relationshipUri,
	}
	// log.Println(data)
	buf, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	// log.Println(string(buf))
	body, err := session.Send(url, string(buf))
	// log.Println(body)
	if err != nil {
		return result, err
	}
	// log.Println(session.StatusCode)
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
