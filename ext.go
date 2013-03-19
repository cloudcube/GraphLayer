package graphdb

//定义结构获取数据，都必须streaming传输，提高性能。
type GraphDbInterface interface {

	// 获取服务根节点
	GetServiceRoot() (*ServiceRootTemplate, error)

	// 查询语句查询
	CypherQueries() (*CypherResultTemplate, error)

	// 从索引中移除条目
	RemoveEntriesFromIndex() error

	// 通过匹配查找节点
	FindNodeByMatch() (map[int]*GraphDataTemplate, error)

	// 通过查询语句查找节点
	FindNodeByQuery() (map[int]*GraphDataTemplate, error)

	// 创建唯一节点
	CreateUniqueNode() (*GraphDataTemplate, error)

	// 获得唯一节点
	GetUniqueNode() (*GraphDataTemplate, error)

	// 创建唯一关系
	CreateUniqueRelationship() (*GraphDataTemplate, error)

	// 获得唯一关系
	GetUniqueRelationship() (*GraphDataTemplate, error)

	// 通过匹配从自动索引中查找节点
	FindNodeFromAutomaticIndexByMatch() (map[int]*GraphDataTemplate, error)

	// 通过查询语句从自动索引中查找节点
	FindNodeFromAutomaticIndexByQuery() (map[int]*GraphDataTemplate, error)

	// 为节点创建自动索引
	CreateAutoIndexForNodes() (*IndexTemplate, error)

	// 为关系创建自动索引
	CreateAutoIndexForRelationships() (*IndexTemplate, error)

	// 为节点上的自动索引获取当前状态
	GetNowStatusForAutoIndexing() (bool, error)

	// 列出自动索引的属性
	ListPropertiesAutoIndexed() ([]string, error)

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
