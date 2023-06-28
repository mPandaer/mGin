package gee

import "strings"

const pathSep = "/"

type node struct {
	pattern  string  //待匹配的路由，只有叶子节点不为空值
	part     string  //路由中的一部分
	children []*node //该节点的子节点
	isWild   bool    //是否是模糊匹配
}

// searchChild 在当前节点的子节点中寻找是否匹配part的子节点 只寻找第一个
func (n *node) searchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// searchChildren 在当前节点的子节点中寻找是否匹配part的子节点 满足条件的都要
func (n *node) searchChildren(part string) []*node {
	var children []*node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}
	return children
}

//插入 指定的路由切片
/**
insert
	parts 以/为分割的路由切片
	height 当前匹配的深度
*/
func (n *node) insert(parts []string, height int) {
	pattern := pathSep + strings.Join(parts, pathSep)

	//当匹配深度 >= parts 时，说明匹配完成
	if height >= len(parts) {
		n.pattern = pattern //因为是插入，所以需要记录路由全路径方便读取
		return
	}

	//插入之前 判断是否之前插入过
	part := parts[height]
	child := n.searchChild(part)

	//没有插入
	if child == nil {
		child = &node{
			pattern: pattern,
			part:    part,
			isWild:  string(part[0]) == ":" || string(part[0]) == "*",
		}
		n.children = append(n.children, child)
	}

	//插入节点
	child.insert(parts, height+1)
}

// search 搜索指定的路由切片
func (n *node) search(parts []string, height int) *node {

	//搜索深度 >= parts,或者包含通配符 搜索结束
	if height >= len(parts) || strings.HasPrefix(n.part, "*") {
		//如果没有完整的路由路径，搜索失败
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.searchChildren(part)

	//搜索的具体逻辑
	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}
