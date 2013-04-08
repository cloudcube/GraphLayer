package graphdb

import (
	"encoding/json"
	"errors"
	"log"
	// "strconv"
)

// type IndexInterface interface {
// 	//创建节点索引
// 	CreateNodeIndex() error

// 	//删除节点索引
// 	DeleteNodeIndex() error

// 	//列出所有节点索引
// 	ListNodeIndexes() (map[int]*GraphDataTemplate, error)

// 	//添加节点到索引
// 	AddNodeToIndex(key string, value string, id uint64) (error)

// 为节点创建自动索引
//CreateAutoIndexForNodes() (*IndexTemplate, error)

// 为关系创建自动索引
//CreateAutoIndexForRelationships() (*IndexTemplate, error)
// }

func (session *Session) CreateNodeIndex(indexName string) error {
	session.Method = "post"
	url := session.URL + "/index" + "/node/"
	data := map[string]string{}
	data["name"] = indexName
	s, err := json.Marshal(data)
	if err != nil {
		return errors.New("Can't Marsh data")
	}
	_, err = session.Send(url, string(s))
	if err != nil {
		return err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}

func (session *Session) CreateNodeIndexWithConf(indexName string, indexType string, indexProvider string) error {
	session.Method = "post"
	url := session.URL + "/index" + "/node/"
	config := map[string]string{}
	config["type"] = indexType
	config["provider"] = indexProvider
	data := map[string]interface{}{}
	data["name"] = indexName
	data["config"] = config
	s, err := json.Marshal(data)
	if err != nil {
		return errors.New("Can't Marsh data")
	}
	log.Println(data)
	body, err := session.Send(url, string(s))
	if err != nil {
		return err
	}
	log.Println(body)
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}

func (session *Session) ListNodeIndexes() (*GraphDataTemplate, error) {
	session.Method = "get"
	url := session.URL + "/index" + "/node/"
	body, err := session.Send(url, "")
	if err != nil {
		return nil, err
	}
	tmpData, err := session.Unmarshal(body)
	if err != nil {
		return nil, err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return tmpData[0], session.NewError(errorList)
}

func (session *Session) DeleteNodeIndex(key string) error {
	session.Method = "delete"
	url := session.URL + "/index" + "/node"
	_, err := session.Send(url+"/"+key, "")
	if err != nil {
		return err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)

}

func (session *Session) AddNodeToIndex(key string, value string, indexName string, id uint64) (*GraphDataTemplate, error) {
	node, err := session.GetNode(id)
	if err != nil {
		return nil, err
	}
	nodeSelf := node.Self
	url := session.URL + "/index/node/" + indexName
	data := map[string]string{}
	data["key"] = key
	data["value"] = value
	data["uri"] = nodeSelf
	s, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	session.Method = "post"
	body, err := session.Send(url, string(s))
	if err != nil {
		return nil, err
	}
	result, err := session.Unmarshal(body)
	if err != nil {
		return nil, err
	}

	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return result[0], session.NewError(errorList)
}

func (session *Session) CreateAutoIndexForNodes(indexName string, indexType string, indexProvider string) error {
	session.Method = "post"
	config := map[string]string{}
	config["type"] = indexType
	config["provider"] = indexProvider
	data := map[string]interface{}{}
	data["name"] = indexName
	data["config"] = config
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}
	url := session.URL
	url += "/" + "index" + "/" + "node"
	body, err := session.Send(url, string(s))
	if err != nil {
		return err
	}
	log.Println(body)
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}

func (session *Session) CreateAutoIndexForRelationships(indexName string, indexType string, indexProvider string) error {
	session.Method = "post"
	url := session.URL
	url += "/" + "index" + "/" + "relationship"
	config := map[string]string{
		"type":     indexType,
		"provider": indexProvider,
	}
	data := map[string]interface{}{
		"name":   indexName,
		"config": config,
	}
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body, err := session.Send(url, string(s))
	if err != nil {
		return err
	}
	log.Println("body content:")
	log.Println(body)
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}
