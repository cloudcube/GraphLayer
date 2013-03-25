package graphdb

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

/*type PropertyInterface interface {

	// 设置节点属性
	SetPropertyOnNode(nodeId uint64, data map[string]string) error

	// 更新节点属性
	UpdateNodeProperties(nodeId uint64,data map[string]string) error

	// 获取节点属性
	GetPropertiesForNode(nodeId uint64) (*GraphDataTemplate, error)

	// 删除节点属性
	DeletePropertyFromNode(nodeId uint64, key string) error

	//删除所有节点属性
	DeletePropertiesFromNode(nodeId uint64) error

	// 更新关系属性
	UpdateRelationshipProperties() error

	// 获取关系属性
	GetPropertiesFromRelationship() (*GraphDataTemplate, error)

	// 移除关系属性
	RemovePropertiesFromRelationship() error
}*/

// 设置节点属性
func (session *Session) SetPropertyOnNode(nodeId uint64, data map[string]string) error {
	node, err := session.GetNode(nodeId)
	if err != nil {
		return err
	}
	session.Method = "put"
	if err != nil {
		return err
	}
	for k, v := range data {
		k = strings.TrimSpace(k)
		_, err = session.Send(node.Properties+"/"+k, strconv.Quote(v))
		if err != nil {
			return err
		}
	}
	errorList := map[int]error{
		404: errors.New("Node not found."),
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}

// 更新节点属性,会替换掉原来的所有属性.
func (session *Session) UpdateNodeProperties(nodeId uint64, data map[string]string) error {
	node, err := session.GetNode(nodeId)
	if err != nil {
		return err
	}
	session.Method = "put"
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = session.Send(node.Properties, string(s))
	if err != nil {
		return err
	}
	errorList := map[int]error{
		404: errors.New("Node not found."),
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}

// 获取节点属性
func (session *Session) GetPropertiesForNode(nodeId uint64) (nodeData *GraphDataTemplate, err error) {
	node, err := session.GetNode(nodeId) // find properties for node
	if err != nil {
		return nodeData, err
	}
	session.Method = "get"
	body, err := session.Send(node.Properties, "")
	if err != nil {
		return nodeData, err
	}
	// pack json string into variable "data" so the json unmarshaler knows where to put it on struct type NeoTemplate
	jsonData, err := session.Pack("data", body)
	if err != nil {
		return nodeData, err
	}
	//convert json -> string and unmarshal -> NeoTemplate
	template, err := session.Unmarshal(string(jsonData))
	if err != nil {
		return nodeData, err
	}
	errorList := map[int]error{
		404: errors.New("Node or Property not found."),
		204: errors.New("No properties found."),
	}
	return template[0], this.NewError(errorList)
}

// 删除节点属性
func (session *Session) DeletePropertyFromNode(nodeId uint64, key string) error {
	node, err := session.GetNode(nodeId) // find properties for node
	if err != nil {
		return err
	}
	session.Method = "delete"
	_, err = session.Send(node.Properties+"/"+string(key), "")
	if err != nil {
		return err
	}
	errorList := map[int]error{
		404: errors.New("Node or Property not found."),
	}
	return session.NewError(errorList)
}

//删除所有节点属性
func (session *Session) DeletePropertiesFromNode(nodeId uint64) error {
	node, err := session.GetNode(nodeId)
	if err != nil {
		return err
	}
	session.Method = "delete"
	_, err = session.Send(node.Properties, "")
	if err != nil {
		return err
	}

	errorList := map[int]error{
		404: errors.New("Node or Property not found."),
	}
	return session.NewError(errorList)
}

// 更新关系属性
func (session *Session) UpdateRelationshipProperties(relId uint64, data map[string]string) error {
	relationship, err := session.GetRelationshipById(relId)
	if err != nil {
		return err
	}
	session.Method = "put"
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = session.Send(relationship[0].ID, string(buf))
	if err != nil {
		return err
	}
	errorList := map[int]error{
		404: errors.New("Node or Property not found."),
	}
	return session.NewError(errorList)
}
