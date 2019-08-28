package graph

/** 邻接表 数据结构表示 **/
// 边
type EdgeNode struct {
	adjIndex int         // 在adjList中的索引位置
	info     interface{} //边间的附属信息，包括权值等
	next     *EdgeNode   // 下一个节点
}

// 顶点表结点
type VertexNode struct {
	data interface{} // 顶点的值描述
	edge *EdgeNode   //顶点指向的边
}

// 邻接表
type GraphAdjList struct {
	nodes     []*VertexNode // 节点数据
	vertexNum int           // 顶点数
	edgesNum  int           // 边数
	kind      int           // 图的种类 0:无向图 1:有向图
}
