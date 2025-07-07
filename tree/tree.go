package tree

type BTreeNode[T any] struct {
	Data  T
	Left  *BTreeNode[T]
	Right *BTreeNode[T]
}

func (n *BTreeNode[T]) appendLeftNode(node *BTreeNode[T]) {
	n.Left = node
}

func (n *BTreeNode[T]) appendRightNode(node *BTreeNode[T]) {
	n.Right = node
}

type MTreeNode[T any] struct {
	Data     T
	Children []*MTreeNode[T]
}

func (n *MTreeNode[T]) appendChild(childNode *MTreeNode[T]) {
	if n.Children == nil {
		n.Children = []*MTreeNode[T]{}
	}
	n.Children = append(n.Children, childNode)
}
