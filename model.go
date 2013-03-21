package graphdb

import ()

type ServiceRootTemplate struct {
	Extensions        []string
	Node              string
	ReferenceNode     string
	RelationshipIndex string
	ExtensionsInfo    string
	RelationshipTypes string
	Batch             string
	Cypher            string
	Neo4jVersion      string
}

type GraphDataTemplate struct {
	ID                  uint64
	Relationships       string
	RelationshipsOut    string
	RelationshipsIn     string
	RelationshipsAll    string
	RelationshipsCreate string
	Data                map[string]interface{}
	Traverse            string
	Property            string
	Properties          string
	Self                string
	Extensions          map[string]interface{}
	Start               string        // relationships & traverse // returns both obj & string
	End                 string        // relationships & traverse // returns both obj & string
	Type                string        // relationships & traverse
	Indexed             string        // index related
	Length              string        // traverse framework
	Nodes               []interface{} // traverse framework
	TRelationships      []interface{} // traverse framework

}

type CypherResultTemplate struct {
}

type IndexTemplate struct {
	Template string
	Type     string
	Provider string
}

type Session struct {
	Method     string // which http method
	StatusCode int    // last http status code received
	URL        string
	Username   string
	Password   string
}

type Error struct {
	List map[int]error
	Code int
}
