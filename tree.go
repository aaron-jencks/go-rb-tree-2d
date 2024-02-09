package gorbtree2d

type treeNode interface {
	value() interface{}
	right() treeNode
	left() treeNode
	parent() treeNode
	attach_right(n treeNode)
	attach_left(n treeNode)
	attach_parent(n treeNode)
}

type RedBlackTree struct {
	root  treeNode
	count int
}

func (t RedBlackTree) size() int {
	return t.count
}

func (t *RedBlackTree) insert(v interface{}) {

}

func (t *RedBlackTree) delete(v interface{}) {

}
