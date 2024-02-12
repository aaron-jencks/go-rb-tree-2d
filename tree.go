package gorbtree2d

type RB_COLORS int

const (
	RB_RED RB_COLORS = iota
	RB_BLACK
)

// a function that returns true if the first element is less than the second
type lessFunc[T any] func(T, T) bool

// represents a node to be used in a tree or graph
type treeNode[T any] interface {
	Value() T                    // returns the value of the node
	Right() treeNode[T]          // returns the right child of the node, or nil if there isn't one
	Left() treeNode[T]           // returns the left child of the node, or nil if there isn't one
	Parent() treeNode[T]         // returns the parent of the node, or nil if there isn't one
	Attach_right(n treeNode[T])  // attaches a child node to the right of this node, updates the parent of the child node
	Attach_left(n treeNode[T])   // attaches a child node to the left of this node, updates the parent of the child node
	Attach_parent(n treeNode[T]) // attaches a parent node to this node
	Detach_right() treeNode[T]   // removes the right child of this node and returns it, updates the parent of the child if it's not nil
	Detach_left() treeNode[T]    // removes the left child of this node and returns it, updates the parent of the child if it's not nil
	Detach_parent() treeNode[T]  // removes the parent of this node and returns it
	Get_color() RB_COLORS        // returns the current color of this node
	Set_color(c RB_COLORS)       // sets the color of the current node
}

type RedBlackNode[T any] struct {
	v T
	r treeNode[T]
	l treeNode[T]
	p treeNode[T]
	c RB_COLORS
}

func CreateNode[T any](value T) *RedBlackNode[T] {
	return &RedBlackNode[T]{
		v: value,
	}
}

func (n *RedBlackNode[T]) Get_color() RB_COLORS {
	return n.c
}

func (n *RedBlackNode[T]) Set_color(c RB_COLORS) {
	n.c = c
}

func (n *RedBlackNode[T]) Value() T {
	return n.v
}

func (n *RedBlackNode[T]) Right() treeNode[T] {
	return n.r
}

func (n *RedBlackNode[T]) Left() treeNode[T] {
	return n.l
}

func (n *RedBlackNode[T]) Parent() treeNode[T] {
	return n.p
}

func (n *RedBlackNode[T]) Attach_right(r treeNode[T]) {
	if r != nil {
		r.Attach_parent(n)
	}
	n.r = r
}

func (n *RedBlackNode[T]) Attach_left(l treeNode[T]) {
	if l != nil {
		l.Attach_parent(n)
	}
	n.l = l
}

func (n *RedBlackNode[T]) Attach_parent(p treeNode[T]) {
	n.p = p
}

func (n *RedBlackNode[T]) Detach_right() treeNode[T] {
	result := n.r
	n.r = nil
	if result != nil {
		result.Detach_parent()
	}
	return result
}

func (n *RedBlackNode[T]) Detach_left() treeNode[T] {
	result := n.l
	n.l = nil
	if result != nil {
		result.Detach_parent()
	}
	return result
}

func (n *RedBlackNode[T]) Detach_parent() treeNode[T] {
	result := n.p
	n.p = nil
	return result
}

type RedBlackTree[T any] struct {
	less  lessFunc[T]
	root  treeNode[T]
	count int
}

func CreateTree[T any](less lessFunc[T]) *RedBlackTree[T] {
	return &RedBlackTree[T]{
		less: less,
	}
}

func (t RedBlackTree[T]) Size() int {
	return t.count
}

func (t *RedBlackTree[T]) RotateNodeRight(n treeNode[T]) {
	p := n.Parent()
	l := n.Left()
	if l != nil {
		if l.Right() != nil {
			n.Attach_left(l.Right())
		}
		l.Attach_right(n)
		l.Attach_parent(p)
		if p == nil {
			t.root = l
		}
	}
}

func (t *RedBlackTree[T]) RotateNodeLeft(n treeNode[T]) {
	p := n.Parent()
	r := n.Right()
	if r != nil {
		if r.Left() != nil {
			n.Attach_right(r.Left())
		}
		r.Attach_left(n)
		r.Attach_parent(p)
		if p == nil {
			t.root = r
		}
	}
}

func (t *RedBlackTree[T]) Insert(v T) {
	if t.Size() == 0 {
		t.root = CreateNode[T](v)
		t.count++
		return
	}

	n := CreateNode[T](v)

	current := t.root
	for {
		if t.less(v, current.Value()) {
			if current.Left() != nil {
				current = current.Left()
				continue
			}
			current.Attach_left(n)
			break
		}

		if current.Right() != nil {
			current = current.Right()
			continue
		}
		current.Attach_right(n)
		break
	}

	// TODO perform red black corrections
	for n.Get_color() == RB_RED {

	}
}

func (t *RedBlackTree[T]) Delete(v T) {

}
