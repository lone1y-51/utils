package tree

type BRTree struct {
	Root *BRTreeNode
}

func (tr *BRTree) FindNode(dst int) (node *BRTreeNode, err error) {
	if tr.Root == nil {
		return nil, ErrEmptyTree
	}
	curNode := tr.Root
	for curNode != nil {
		if dst == curNode.Value {
			return curNode, nil
		} else if dst > curNode.Value {
			curNode = curNode.RChildNode
			continue
		} else if dst < curNode.Value {
			curNode = curNode.LChildNode
			continue
		}
	}
	return nil, ErrNotFound
}

func (tr *BRTree) Insert(value int) error {
	newNode := &BRTreeNode{
		Value: value,
		Color: Red,
	}
	// 根节点为空
	if tr.Root == nil {
		newNode.ChangeToBlackNode()
		tr.Root = newNode
		return nil
	}
	parentNode := tr.findInsertedParentNode(newNode)
	// 父节点的值和新节点的值相同
	if parentNode.Value == newNode.Value {
		return nil
	}
	newNode.ParentNode = parentNode
	return tr.insertNode(newNode, parentNode)
}

func (tr *BRTree) findInsertedParentNode(node *BRTreeNode) *BRTreeNode {
	parentNode := tr.Root
	for {
		if node.Value == parentNode.Value {
			break
		} else if node.Value > parentNode.Value {
			if parentNode.RChildNode == nil {
				break
			}
			parentNode = parentNode.RChildNode
			continue
		} else if node.Value < parentNode.Value {
			if parentNode.LChildNode == nil {
				break
			}
			parentNode = parentNode.LChildNode
			continue
		}
	}
	return parentNode
}

func (tr *BRTree) insertNode(node, parentNode *BRTreeNode) error {
	// 表示是根节点
	if parentNode == nil {
		node.ChangeToBlackNode()
		return nil
	}
	if node.Value < parentNode.Value {
		node.ParentNode = parentNode
		parentNode.LChildNode = node
		// 父节点是黑色
		if !parentNode.IsRedNode() {
			return nil
		}
		uncleNode := parentNode.FindBrotherNode()
		if uncleNode == nil || !uncleNode.IsRedNode() {
			// 叔叔节点不存在或者是黑色
			parentNode.ChangeToBlackNode()
			parentNode.ParentNode.ChangeToRedNode()
			if parentNode == parentNode.ParentNode.LChildNode {
				parentNode.ParentNode.RightHand()
			} else {
				parentNode.ParentNode.LeftHand()
			}
		} else {
			parentNode.ChangeToBlackNode()
			uncleNode.ChangeToBlackNode()
			parentNode.ParentNode.ChangeToRedNode()
			return tr.insertNode(parentNode.ParentNode, parentNode.ParentNode.ParentNode)
		}
	} else if node.Value > parentNode.Value {
		node.ParentNode = parentNode
		parentNode.RChildNode = node
		if !parentNode.IsRedNode() {
			return nil
		}
		uncleNode := parentNode.FindBrotherNode()
		if uncleNode == nil || !uncleNode.IsRedNode() {
			parentNode.ChangeToBlackNode()
			parentNode.ParentNode.ChangeToRedNode()
			if parentNode == parentNode.ParentNode.LChildNode {
				parentNode.ParentNode.RightHand()
			} else {
				parentNode.ParentNode.LeftHand()
			}
		} else {
			parentNode.ChangeToBlackNode()
			uncleNode.ChangeToBlackNode()
			parentNode.ParentNode.ChangeToRedNode()
			return tr.insertNode(parentNode.ParentNode, parentNode.ParentNode.ParentNode)
		}
	}
	return nil
}
