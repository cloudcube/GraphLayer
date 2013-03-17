package graphdb

//定义结构获取数据，都必须streaming传输，提高性能。
type GraphDbInterface interface {

	// 获取服务根节点
	GetServiceRoot() error

	// 查询语句查询
	CypherQueries() error

	// 从索引中移除条目
	RemoveEntriesFromIndex() error

	// 通过匹配查找节点
	FindNodeByMatch()

	// 通过查询语句查找节点
	FindNodeByQuery()

	// 创建唯一节点
	CreateUniqueNode() error

	// 获得唯一节点
	GetUniqueNode() error

	// 创建唯一关系
	CreateUniqueRelationship() error

	// 获得唯一关系
	GetUniqueRelationship() error

	// 通过匹配从自动索引中查找节点
	FindNodeFromAutomaticIndexByMatch() error

	// 通过查询语句从自动索引中查找节点
	FindNodeFromAutomaticIndexByQuery() error

	// 为节点创建自动索引
	CreateAutoIndexForNodes() error

	// 为关系创建自动索引
	CreateAutoIndexForRelationships() error

	// 为节点上的自动索引获取当前状态
	GetNowStatusForAutoIndexing() error

	// 列出自动索引的属性
	ListPropertiesAutoIndexed() error

	// 为自动索引添加属性
	AddPropertyForAutoIndexing() error

	// 为自动索引添加属性
	RemovePropertyForAutoIndexing() error

	// 使用一个返回过滤条件进行遍历
	TraversalByFilter() error

	// 从一个遍历返回关系
	GetRelationshipsFromTraversal() error

	// 从一个遍历返回路径
	GetPathsFromTraversal() error

	// 遍历返回低于一定深度的节点
	GetNodesBelowDepthAtTraversal() error

	// 创建一个分页的遍历
	CreatedPagedTraversers() error

	// 获取分页遍历的分页结果
	GetPagingResultPagedTraverser() error

	// 为分页索引设置分页大小
	SetPageSizeForPagedTraverser() error

	// 设置分页遍历的超时时间
	SetPagedTraverserTimeout() error

	// 寻找最短路径
	FindShortestPaths() error

	// 在关系上执行Dijkstra
	ExecDijkstraOnRelationships() error
}
