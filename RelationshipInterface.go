package graphdb

type RelationshipInterface interface {

	//创建关系
	CreateRelationship() error

	//根据ID查询关系
	GetRelationshipById() error

	//获得所有关系
	GetRelationships() error

	//获得关系类型
	GetRelationshipTypes() error

	//删除关系
	DeleteRelationship() error
}
