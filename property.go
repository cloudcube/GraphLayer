package graphdb

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

/*type PropertyInterface interface {

	// 根据名称查找属性
	GetPropertyByName(nodeId uint64, name string) (string, error)

	// 设置节点属性
	SetPropertyOnNode(nodeId uint64, data map[string]string, isReplace bool) error

	// 更新节点属性
	UpdateNodeProperties() error

	// 获取节点属性
	GetPropertiesFromNode() (*GraphDataTemplate, error)

	// 删除节点属性
	DeletePropertyFromNode() error

	// 设置关系属性
	SetRelationshipProperties() error

	// 更新关系属性
	UpdateRelationshipProperties() error

	// 获取关系属性
	GetPropertiesFromRelationship() (*GraphDataTemplate, error)

	// 移除关系属性
	RemovePropertiesFromRelationship() error
}*/

// 根据名称查找属性
func (session *Session) GetPropertyByName(nodeId uint64, name string) (string, error) {
	if len(name) < 1 {
		return "", errors.New("Property name must be at least 1 character.")
	}
	node, err := session.GetNode(nodeId)
	if err != nil {
		return "", err
	}
	session.Method = "get"
	body, err := session.Send(node.Property+"/"+name, "")
	if err != nil {
		return "", err
	}
	errorList := map[int]error{
		404: errors.New("Node or Property not found."),
		204: errors.New("No Properties found."),
	}
	return body, session.NewError(errorList)
}

// 设置节点属性
func (session *Session) SetPropertyOnNode(nodeId uint64, data map[string]string, isReplace bool) error {
	node, err := session.GetNode(nodeId)
	if err != nil {
		return err
	}
	session.Method = "put"
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if isReplace {
		_, err := session.Send(node.Properties, string(s))
		if err != nil {
			return err
		}
	} else {
		for k, v := range data {
			k = strings.TrimSpace(k)
			_, err := session.Send(node.Properties+"/"+k, strconv.Quote(v))
			if err != nil {
				return err
			}
		}
	}
	errorList := map[int]error{
		404: errors.New("Node not found."),
		400: errors.New("Invalid data send."),
	}
	return session.NewError(errorList)
}
