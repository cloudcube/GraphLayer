package GraphLayer

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// type RelationshipInterface interface {

// 	//创建关系
// 	CreateRelationship(src uint64, dest uint64, relDesc map[string]string, relType string) (map[int]*GraphDataTemplate, error)

// 	//根据ID查询关系
// 	GetRelationshipById(relId uint64) (map[int]*GraphDataTemplate, error)

// 	//获得关系,包含关系方向和是否标记
// 	GetRelationshipsOnNode(nodeId uint64, typed string, direction string) (map[int]*GraphDataTemplate, error)

// 	//获得关系类型
// 	GetRelationshipTypes() ([]string, error)

// 	//删除关系
// 	DeleteRelationship(relId uint64) error
// }

//创建关系 source node id uint64,destination node id uint64, relationship description map[string]string,relation type string,return any errors raise error
func (session *Session) CreateRelationship(src uint64, dest uint64, relDesc map[string]string, relType string) (dataSet map[int]*GraphDataTemplate, err error) {
	dataSet = make(map[int]*GraphDataTemplate)
	var destNode, srcNode *GraphDataTemplate
	destNode, err = session.GetNode(dest) // find properties for destination node so we can tie it into the relationship
	if err != nil {
		return
	}
	srcNode, err = session.GetNode(src) // find properties for src node..
	if err != nil {
		return
	}
	j := map[string]interface{}{} // empty map: keys are always strings in json, values vary
	j["to"] = destNode.Self
	j["type"] = relType             // type of relationship
	j["data"] = map[string]string{} // empty array
	j["data"] = relDesc             // add data to relationship
	s, err := json.Marshal(j)
	if err != nil {
		err = errors.New("Unable to Marshal Json data")
		return
	}
	session.Method = "post"
	var body string
	body, err = session.Send(srcNode.RelationshipsCreate, string(s)) // srcNode.RelationshipsCreate actually contains the full URL
	if err != nil {
		return
	}
	dataSet, err = session.Unmarshal(body)
	if err != nil {
		return
	}
	errorList := map[int]error{
		404: errors.New("Node or 'to' node not found."),
		400: errors.New("Invalid data sent."),
	}
	return dataSet, session.NewError(errorList)
}

//根据ID查询关系
func (session *Session) GetRelationshipById(relId uint64) (dataSet map[int]*GraphDataTemplate, err error) {
	if relId < 0 {
		return dataSet, errors.New("Invalid Relationship id specified.")
	}
	session.Method = "get"
	// url := session.URL + "/relationship/"
	url := session.URL + "/relationship/"

	body, err := session.Send(url+strconv.FormatUint(uint64(relId), 10), "")
	if err != nil {
		return dataSet, err
	}
	dataSet, err = session.Unmarshal(body)
	if err != nil {
		return nil, err
	}
	errorList := map[int]error{
		404: errors.New("Relatoinship not found."),
	}
	return dataSet, session.NewError(errorList)
}

//获得关系类型
func (session *Session) GetRelationshipTypes() ([]string, error) {
	var types []string
	session.Method = "get"
	url := session.URL + "/relationship/types"
	//data := map[string]string{}
	body, err := session.Send(url, "")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &types)
	return types, nil
}

//Get All Relationships By node id
func (session *Session) GetRelationshipsOnNode(nodeId uint64, typed string, direction string) (dataSet map[int]*GraphDataTemplate, err error) {
	node, err := session.GetNode(nodeId)
	if err != nil {
		return nil, err
	}
	session.Method = "get"
	direction = strings.ToLower(direction)
	url := ""
	switch direction {
	case "in":
		url = node.RelationshipsIn
	case "out":
		url = node.RelationshipsOut
	case "all":
		fallthrough
	default:
		url = node.RelationshipsAll
	}

	body, err := session.Send(url+"/"+typed, "")
	if err != nil {
		return nil, err
	}

	template, err := session.Unmarshal(body)
	if err != nil {
		return nil, err
	}

	errorList := map[int]error{
		404: errors.New("Node not found."),
	}
	return template, session.NewError(errorList)
}

//删除关系
func (session *Session) DeleteRelationship(relId uint64) error {
	session.Method = "delete"
	url := session.URL + "/relationship/"
	_, err := session.Send(url+strconv.FormatUint(uint64(relId), 10), "")
	if err != nil {
		return err
	}
	errorList := map[int]error{
		404: errors.New("Relationship not found!"),
	}
	return session.NewError(errorList)
}
