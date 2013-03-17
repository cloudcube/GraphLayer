package graphdb

type NodeInterface interface {
	// 创建节点
	CreateNode() error

	// 获取节点
	GetNode() error

	// 删除节点
	DeleteNode() error
}
