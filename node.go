package graphdb

type NodeInterface interface {
	// 创建节点
	CreateNode() (*GraphDataTemplate, error)

	// 获取节点
	GetNode() (*GraphDataTemplate, error)

	// 删除节点
	DeleteNode() error
}
