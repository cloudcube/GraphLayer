/*
Traversals are performed from a start node. The traversal is controlled by the URI and the body sent
with the request.
returnType
The kind of objects in the response is determined by traverse/{returnType} in the URL. returnType
can have one of these values:
• node
• relationship
• path: contains full representations of start and end node, the rest are URIs.
• fullpath: contains full representations of all nodes and relationships.
To decide how the graph should be traversed you can use these parameters in the request body:
order
Decides in which order to visit nodes. Possible values:
• breadth_first: see Breadth-first search <http://en.wikipedia.org/wiki/Breadth-first_search>.
• depth_first: see Depth-first search <http://en.wikipedia.org/wiki/Depth-first_search>
relationships
Decides which relationship types and directions should be followed. The direction can be one of:
• all
• in
• out
uniqueness
Decides how uniqueness should be calculated. For details on different uniqueness values see the
Java API on Uniqueness <http://components.neo4j.org/neo4j/1.8.2/apidocs/org/neo4j/kernel/
Uniqueness.html>. Possible values:
• node_global
• none
• relationship_global
• node_path
• relationship_path
prune_evaluator
Decides whether the traverser should continue down that path or if it should be pruned so that
the traverser won’t continue down that path. You can write your own prune evaluator as (see
Section 18.13.1, “Traversal using a return filter” or use the built-in none prune evaluator.
304REST API
return_filter
Decides whether the current position should be included in the result. You can provide your own
code for this (see Section 18.13.1, “Traversal using a return filter”), or use one of the built-in
filters:
• all
• all_but_start_node
max_depth
Is a short-hand way of specifying a prune evaluator which prunes after a certain depth. If not
specified a max depth of 1 is used and if a prune_evaluator is specified instead of a max_depth, no
max depth limit is set.
The position object in the body of the return_filter and prune_evaluator is a Path <http://
components.neo4j.org/neo4j/1.8.2/apidocs/org/neo4j/graphdb/Path.html> object representing the path
from the start node to the current traversal position.
*/

package graphdb

import (
	"encoding/json"
	"errors"
	"strconv"
)

func (session *Session) Traversal(startNode uint64, order string, return_filter map[string]string, prune_evaluator map[string]string, uniqueness string, relationships []map[string]string, max_depth uint64) (dataResults map[int]*GraphDataTemplate, err error) {
	session.Method = "post"
	url := session.URL
	url += "/" + "node"
	if startNode >= 0 {
		url += "/" + strconv.FormatUint(startNode, 10)
	}
	url += "/" + "traverse" + "/" + "node"
	data := map[string]interface{}{}
	if len(order) > 0 {
		data["order"] = order
	}
	if return_filter != nil {
		data["return_filter"] = return_filter
	}
	if prune_evaluator != nil {
		data["prune_evaluator"] = prune_evaluator
	}
	if len(uniqueness) == 0 {
		return dataResults, errors.New("Uniqueness has no value")
	}
	data["uniqueness"] = uniqueness
	if len(relationships) == 0 {
		return dataResults, errors.New("Relationship is empty")
	}
	data["relationships"] = relationships
	if max_depth == 0 {
		return dataResults, errors.New("You'd better associate a value for max_depth")
	}
	data["max_depth"] = max_depth
	buf, err := json.Marshal(data)
	if err != nil {
		return
	}
	body, err := session.Send(url, string(buf))
	if err != nil {
		return
	}
	dataResults, err = session.Unmarshal(body)
	if err != nil {
		return
	}
	errorList := map[int]error{
		400: errors.New("Invalid data send."),
	}
	return dataResults, session.NewError(errorList)
}
