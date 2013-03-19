package graphdb

type RelationshipInterface interface {

	//创建关系
	CreateRelationship() (*map[int]*GraphDataTemplate, error)

	//根据ID查询关系
	GetRelationshipsById() (map[int]*GraphDataTemplate, error)

	//获得所有关系
	GetRelationships() (map[int]*GraphDataTemplate, error)

	//获得关系类型
	GetRelationshipTypes() ([]string, error)

	//删除关系
	DeleteRelationship() error
}
