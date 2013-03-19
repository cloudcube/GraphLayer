package graphdb

import (
	neo4j "neo4j"
)

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
	*neo4j.NeoTemplate
}

type CypherResultTemplate struct {
	*neo4j.ResultSet
}

type IndexTemplate struct {
	Template string
	Type     string
	Provider string
}
