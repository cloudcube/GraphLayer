package graphdb

type PropertyInterface interface {

	// 根据名称查找属性
	GetPropertyByName() (string, error)

	// 设置节点属性
	SetNodeProperties() error

	// 更新节点属性
	UpdateNodeProperties() error

	// 获取节点属性
	GetPropertiesFromNode() (*GraphDataTemplate, error)

	// 删除节点属性
	DeletePropertyFromNode() error

	// 设置关系属性
	SetRelationshipProperties() error

	// 更新关系属性
	UpdateRelationshipProperties() error

	// 获取关系属性
	GetPropertiesFromRelationship() (*GraphDataTemplate, error)

	// 移除关系属性
	RemovePropertiesFromRelationship() error
}
