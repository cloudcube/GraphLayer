package graphdb

import (
	"encoding/json"
	"errors"
	// "strings"
	// "log"
	"strconv"
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
	return session.findNodeFromAutomaticIndexByUri(url)
}

func (session *Session) FindNodeFromAutomaticIndexByQuery(nodeKey string, nodeValue string) (dataResults map[int]*GraphDataTemplate, err error) {
	session.Method = "get"
	url := session.URL
	url += "/" + "index" + "/" + "auto" + "/" + "node"
	url += "/" + "?" + "query="
	if len(nodeKey) == 0 {
		return dataResults, errors.New("nodekey is nil!")
	}
	url += nodeKey
	if len(nodeValue) == 0 {
		return dataResults, errors.New("nodeValue is nil!")
	}
	url += ":" + nodeValue
	return session.findNodeFromAutomaticIndexByUri(url)

}

func (session *Session) findNodeFromAutomaticIndexByUri(uri string) (dataResults map[int]*GraphDataTemplate, err error) {
	body, err := session.Send(uri, "")
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

func (session *Session) CrateAutoIndexWithConf(indexType string, indexProvider string, category string) (err error) {
	session.Method = "post"
	url := session.URL
	indexName := ""
	if category == "node" {
		indexName = "node_auto_index"
		url += "/" + "index" + "/" + "node" + "/"
	}
	if category == "relationship" {
		indexName = "relationship_auto_index"
		url += "/" + "index" + "/" + "relationship" + "/"
	}
	config := map[string]string{
		"type":     indexType,
		"provider": indexProvider,
	}
	postData := map[string]interface{}{
		"name":   indexName,
		"config": config,
	}
	buf, err := json.Marshal(postData)
	if err != nil {
		return err
	}
	_, err = session.Send(url, string(buf))
	if err != nil {
		return err
	}
	return err
}

func (session *Session) dealAutoIndexStatus(category string) string {
	url := session.URL
	if category == "node" {
		url += "/" + "index" + "/" + "auto" + "/" + "node" + "/"
	}
	if category == "relationship" {
		url += "/" + "index" + "/" + "auto" + "/" + "relationship" + "/"
	}
	url += "status"
	return url
}

func (session *Session) GetAutoIndexStatus(category string) (bool, error) {
	session.Method = "get"
	url := session.dealAutoIndexStatus(category)
	body, err := session.Send(url, "")
	if err != nil {
		return false, err
	}
	status, err := strconv.ParseBool(body)
	if err != nil {
		return false, err
	}
	return status, nil
}

func (session *Session) EnableAutoindex(status bool, category string) (err error) {
	session.Method = "put"
	url := session.dealAutoIndexStatus(category)
	buf, err := json.Marshal(status)
	if err != nil {
		return err
	}
	_, err = session.Send(url, string(buf))
	if err != nil {
		return err
	}
	return err
}

func (session *Session) getAutoIndexUri(category string) string {
	url := session.URL
	if category == "node" {
		url += "/" + "index" + "/" + "auto" + "/" + "node" + "/"
	}
	if category == "relationship" {
		url += "/" + "index" + "/" + "auto" + "/" + "relationship" + "/"
	}
	return url
}

// func (session *Session) LookuplistAutoIndexProperties() (indexProperties []string, err error) {
// 	session.Method = "get"
// 	url := session.URL

// }
