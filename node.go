package graphdb

import (
	"encoding/json"
	"errors"
	"log"
)

// type NodeInterface interface {
// 	// 创建节点
// 	AddNode(data map[string]string) (*GraphDataTemplate, error)

// 	// 获取节点
// 	GetNode() (*GraphDataTemplate, error)

// 	// 删除节点
// 	DeleteNode() error
// }

func (session *Session) CreateNode(data map[string]string) (nodeTempalte *GraphDataTemplate, err error) {
	s, err := json.Marshal(data)
	if err != nil {
		return nodeTempalte, errors.New("Unable to Marshal Json data")
	}
	session.Method = "post"
	url := session.URL + "/node"
	body, err := session.Send(url, string(s))
	if err != nil {
		return nodeTempalte, err
	}
	log.Println(body)
	template, err := session.Unmarshal(body) // json.Unmarshal wrapper with some type assertions etc
	if err != nil {
		return nodeTempalte, err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data sent."),
	}
	return template[0], session.NewError(errorList)
}
