package GraphLayer

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
)

var settingFile = "databaseConfig.json"

func (session *Session) Unmarshal(s string) (dataSet map[int]*GraphDataTemplate, err error) {
	var (
		templateNode map[string]interface{}   // blank interface for json.Unmarshal; used for node lvl data
		templateSet  []map[string]interface{} // array of blank interfaces for json.Unmarshal
	)
	dataSet = make(map[int]*GraphDataTemplate)     // make it ready for elements
	err = json.Unmarshal([]byte(s), &templateNode) // unmarshal json data into blank interface. the json pkg will populate with the proper data types
	if err != nil {                                // fails on multiple results
		err = json.Unmarshal([]byte(s), &templateSet) // if unable to unmarshal into single template, try an array of templates instead. If that fails, raise an error
		if err != nil {
			return nil, err
		}
		for _, v := range templateSet {
			data, err := session.unmarshalNode(v) // append NeoTemplate into the data set
			if err != nil {
				return nil, err
			}
			dataSet[len(dataSet)] = data // new array element containing data
		}
	} else {
		template, err := session.unmarshalNode(templateNode)
		if err != nil {
			return nil, err
		}
		dataSet[0] = template // just a single result
	}
	return
}

func (session *Session) NewError(errorList map[int]error) error {
	if errorList != nil {
		errorList[500] = errors.New("Fatal Error 500.") // everything can return a 500 error
	}
	err := &Error{errorList, session.StatusCode}
	return err.CheckError()
}

func (session *Session) unmarshalNode(template map[string]interface{}) (*GraphDataTemplate, error) {
	var (
		data   interface{} // stores data from type assertion
		assert bool        // did the type assertion raise an err?
	)
	node := new(GraphDataTemplate)
	for k, v := range template { // loop result data
		data, assert = v.(map[string]interface{}) // type assertion
		if assert {
			switch vv := data.(type) { // switch on variable type so data/extensions are extracted properly
			case map[string]interface{}:
				switch k {
				case "data":
					node.Data = vv
				case "extensions":
					node.Extensions = vv
				}
			default:
				log.Printf("*Notice: Unknown type in JSON stream: %T from key: %v\n", vv, k)
			}
		} else { // to my knowledge neo4j is only going to pass strings and arrays so if map assertion failed above try an array instead
			data, assert = v.([]interface{}) // normal array?
			if assert {
				switch vv := data.(type) {
				case []interface{}:
					switch k {
					case "nodes":
						node.Nodes = vv
					case "relationships":
						node.TRelationships = vv
					}
				}
			} else { // if nothing else, it must be a string
				data, assert = v.(string)
				if assert {
					// copy string vars into node structure switch on key name
					switch k {
					case "self":
						node.Self, _ = data.(string) // cast it to a string with type assertion
						// "self" provides easy access to the ID property of the node(relationship, index,etc), we'll take advantage and axe it off right now
						selfSlice := strings.Split(string(node.Self), "/")                    // slice string "Self" on each '/' char, -1 gets all instances
						id, atouiErr := strconv.ParseUint(selfSlice[len(selfSlice)-1], 10, 0) // and pull off the last part which is the ID then string -> uint
						if atouiErr != nil {
							return nil, atouiErr
						}
						node.ID = id
					case "traverse":
						node.Traverse, _ = data.(string)
					case "property":
						node.Property, _ = data.(string)
					case "properties":
						node.Properties, _ = data.(string)
					case "outgoing_relationships":
						node.RelationshipsOut, _ = data.(string)
					case "incoming_relationships":
						node.RelationshipsIn, _ = data.(string)
					case "all_relationships":
						node.RelationshipsAll, _ = data.(string)
					case "create_relationship":
						node.RelationshipsCreate, _ = data.(string)
					case "start": // relationships use this
						node.Start, _ = data.(string)
					case "end": // relationships use this
						node.End, _ = data.(string)
					case "type": // relationships use this
						node.Type, _ = data.(string)
					case "length":
						node.Length, _ = data.(string)
					case "indexed": // indices use this
						node.Indexed, _ = data.(string)
					}
				}

			}
		}
	}
	return node, nil
}

func (err *Error) CheckError() error {
	if err.List != nil {
		if err.List[err.Code] != nil {
			return err.List[err.Code]
		}
	}
	return nil
}
