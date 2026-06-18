// package tree
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package tree

type Node interface {
	ID() int64
	PID() int64
	SetChildren([]Node)
}

func ListToTree(rootPid int64, nodes []Node) []Node {
	nodeMap := make(map[int64]Node, len(nodes))
	childrenMap := make(map[int64][]Node)
	for _, node := range nodes {
		if node == nil {
			continue
		}
		nodeMap[node.ID()] = node
	}

	treeList := make([]Node, 0)
	for _, node := range nodes {
		if node == nil {
			continue
		}
		if node.PID() == rootPid {
			treeList = append(treeList, node)
			continue
		}
		if parent, ok := nodeMap[node.PID()]; ok {
			childrenMap[parent.ID()] = append(childrenMap[parent.ID()], node)
		}
	}
	for parentId, children := range childrenMap {
		if parent, ok := nodeMap[parentId]; ok {
			parent.SetChildren(children)
		}
	}
	return treeList
}
