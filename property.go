package graphdb

import (
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
	DeletePropertyFromNode(nodeId uint64) error

	//删除所有节点属性
	DeletePropertiesFromNode(nodeId uint64) error

	// 设置关系属性
	SetRelationshipProperties() error

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
		_, err := session.Send(node.Properties+"/"+k, strconv.Quote(v))
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
