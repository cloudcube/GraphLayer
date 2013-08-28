package goneo4j

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

func (session *Session) GetServiceRoot() (serviceRoot *ServiceRootTemplate, err error) {
	session.Method = "get"
	url := session.URL
	body, err := session.Send(url, "")
	if err != nil {
		return serviceRoot, err
	}
	// log.Println(body)
	err = json.Unmarshal([]byte(body), &serviceRoot)
	if err != nil {
		return serviceRoot, err
	}
	return serviceRoot, nil
}

// 查询语句查询
func (session *Session) CypherQueries(query string, parameters map[string]string) (cypherQueryResult *CypherQueryTemplate, err error) {
	session.Method = "post"
	url := session.URL + "/" + "cypher"
	if len(parameters) == 0 {
		parameters = map[string]string{}
	}
	data := map[string]interface{}{
		"query":  query,
		"params": parameters,
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return cypherQueryResult, err
	}
	body, err := session.Send(url, string(buf))
	if err != nil {
		return cypherQueryResult, err
	}
	log.Println(body)
	err = json.Unmarshal([]byte(body), &cypherQueryResult)
	if err != nil {
		return cypherQueryResult, err
	}
	return cypherQueryResult, nil
}

func (session *Session) RemoveEntriesFromIndex(nodeId uint64, indexName string, indexKey string, indexValue string) error {
	session.Method = "delete"
	url := session.URL
	url += "/" + "index" + "/" + "node"
	if len(indexName) == 0 {
		return errors.New("indexName nil!")
	}
	url += "/" + indexName
	if len(indexKey) > 0 {
		url += "/" + indexKey
	}
	if len(indexValue) > 0 {
		url += "/" + indexValue
	}
	if nodeId == 0 {
		return errors.New("nodeId invalid!")
	}
	url += "/" + strconv.FormatUint(nodeId, 10)
	log.Println(url)
	body, err := session.Send(url, "")
	if err != nil {
		return err
	}
	log.Println("print body ...")
	log.Println(body)
	errorList := map[int]error{
		400: errors.New("Invalid data"),
	}
	return session.NewError(errorList)
}

// 通过匹配查找节点
func (session *Session) FindNodeByMatch(indexName string, indexKey string, indexValue string) (results map[int]*GraphDataTemplate, err error) {
	session.Method = "get"
	url := session.URL
	url += "/" + "index" + "/" + "node"
	if len(indexName) == 0 {
		return results, errors.New("Index Name Invalid.")
	}
	url += "/" + indexName
	if len(indexKey) == 0 {
		return results, errors.New("Index Key Invalid.")
	}
	url += "/" + indexKey
	if len(indexValue) == 0 {
		return results, errors.New("Index Value Invalid.")
	}
	url += "/" + indexValue
	body, err := session.Send(url, "")
	if err != nil {
		return results, err
	}
	results, err = session.Unmarshal(body)
	if err != nil {
		return results, err
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send ."),
	}
	return results, session.NewError(errorList)
}

// 通过查询语句查找节点
func (session *Session) FindNodeByQuery(indexName string, luceneQuery string) (dataResults map[int]*GraphDataTemplate, err error) {
	session.Method = "get"
	url := session.URL + "/" + "index" + "/" + "node"
	if len(indexName) == 0 {
		return dataResults, errors.New("Index can't be nil")
	}
	url += "/" + indexName
	if len(luceneQuery) == 0 {
		return dataResults, errors.New("lucene query can't be nil")
	}
	url += "?" + "query=" + luceneQuery
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
