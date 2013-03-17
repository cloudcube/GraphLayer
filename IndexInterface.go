package graphdb

type IndexInterface interface {
	//创建节点索引
	CreateNodeIndex() error

	//删除节点索引
	DeleteNodeIndex() error

	//列出所有节点索引
	ListNodeIndexes() error

	//添加节点到索引
	AddNodeToIndex() error
}
