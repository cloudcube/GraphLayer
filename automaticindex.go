package graphdb

import (
	"errors"
	// "log"
)

// 通过匹配从自动索引中查找节点
func (session *Session) FindNodeFromAutomaticIndexByMatch(nodeKey string, nodeValue string) (dataResults map[int]*GraphDataTemplate, err error) {
	session.Method = "get"
	url := session.URL
	url += "/" + "index" + "/" + "auto" + "/" + "node"
	if len(nodeKey) == 0 {
		return dataResults, errors.New("nodekey is nil!")
	}
	url += "/" + nodeKey
	if len(nodeValue) == 0 {
		return dataResults, errors.New("nodeValue is nil")
	}
	url += "/" + nodeValue
	body, err := session.Send(url, "")
	if err != nil {
		return dataResults, err
	}
	dataResults, err = session.Unmarshal(body)
	if err != nil {
		return dataResults, err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return dataResults, session.NewError(errorList)
}
