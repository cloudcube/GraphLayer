package graphdb

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

// type NodeInterface interface {
// 	// 创建节点
// 	AddNode(data map[string]string) (*GraphDataTemplate, error)

// 	// 获取节点
// 	GetNode(nodeId uint64) (*GraphDataTemplate, error)

// 	// 删除节点
// 	DeleteNode(nodeId uint64) error
// }

//创建节点
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

// 获取节点
func (session *Session) GetNode(nodeId uint64) (nodeTemplate *GraphDataTemplate, err error) {
	if nodeId < 0 {
		return nodeTemplate, errors.New("Invalid node id specified.")
	}
	session.Method = "get"
	url := session.URL + "/node/"
	body, err := session.Send(url+strconv.FormatUint(uint64(nodeId), 10), "") // convert uint -> string and send http request
	if err != nil {
		return nodeTemplate, err
	}
	template, err := session.Unmarshal(body) // json.Unmarshal wrapper with some type assertions etc
	if err != nil {
		return nodeTemplate, err
	}
	errorList := map[int]error{
		404: errors.New("Node not found."),
	}
	return template[0], session.NewError(errorList)
}

// 删除节点
func (session *Session) DeleteNode(nodeId uint64) error {
	node, err := session.GetNode(nodeId) // find properties for node
	if err != nil {
		return err
	}
	session.Method = "delete"
	_, err = session.Send(node.Self, "")
	if err != nil {
		return err
	}
	errorList := map[int]error{
		404: errors.New("Node not found."),
		409: errors.New("Unable to delete node. May still have relationships."),
	}
	return session.NewError(errorList)
}
