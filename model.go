package graphdb

type ServiceRootTemplate struct {
	Extensions         map[string]interface{}
	Node               string
	Reference_Node     string
	Node_Index         string
	Relationship_Index string
	Extensions_Info    string
	Relationship_Types string
	Batch              string
	Cypher             string
	Neo4j_Version      string
}

//Relationship Template
type RelationshipTemplate struct {
	Extensions map[string]interface{}
	Start      string
	Property   string
	Self       string
	Properties string
	Type       string
	End        string
	Data       map[string]interface{}
}

//Node Template
type NodeTemplate struct {
	Extensions                 map[string]interface{}
	OutgoingRelationships      string
	PagedTraverse              string
	AllTypedRelationships      string
	Property                   string
	AllRelationships           string
	Self                       string
	OutgoingTypedRelationships string
	Properties                 string
	IncomingRelationships      string
	IncomingTypedRelationships string
	CreateRelationship         string
	Data                       map[string]interface{}
}

//Index Template
type IndexTemplate struct {
	Template string
	Provider string
	Type     string
}

type config struct {
	configType string
	provider   string
}

type IndexConfig struct {
	Name string
	config
}

//GraphDataTemplate
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

type CypherQueryTemplate struct {
	Columns []string
	Data    [][]interface{}
}

type Session struct {
	Method     string // which http method
	StatusCode int    // last http status code received
	Location   string //will return a location when creat paged traverser
	URL        string
	Username   string
	Password   string
}

type Error struct {
	List map[int]error
	Code int
}
