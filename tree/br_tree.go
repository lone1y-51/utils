package tree

import "fmt"

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
	fmt.Printf("newNode %s parentNode %s\n", newNode.ToString(), parentNode.ToString())
	return tr.insertNode(newNode, parentNode)
}

func (tr *BRTree) ToString() string {
	var nodeArr []*BRTreeNode
	tr.replaceRoot()
	nodeArr = append(nodeArr, tr.Root)
	index := 0
	for len(nodeArr) != index {
		if nodeArr[index].LChildNode != nil {
			nodeArr = append(nodeArr, nodeArr[index].LChildNode)
		}
		if nodeArr[index].RChildNode != nil {
			nodeArr = append(nodeArr, nodeArr[index].RChildNode)
		}
		index++
	}
	result := ""
	for _, node := range nodeArr {
		result += fmt.Sprintf("%s ", node.ToString())
	}
	return result
}

func (tr *BRTree) replaceRoot() {
	for {
		if tr.Root.IsRoot() {
			return
		}
		tr.Root = tr.Root.ParentNode
	}
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
		// 插入节点是父节点的左子节点
		node.ParentNode = parentNode
		parentNode.LChildNode = node
		// 父节点是黑色
		if !parentNode.IsRedNode() {
			node.ChangeToRedNode()
			node.ParentNode = parentNode
			parentNode.LChildNode = node.ParentNode
			return nil
		}
		uncleNode := parentNode.FindBrotherNode()
		if uncleNode == nil || !uncleNode.IsRedNode() {
			// 叔叔节点不存在或者是黑色
			ppNode := parentNode.ParentNode
			if parentNode == ppNode.LChildNode {
				// 父亲节点是祖父节点的左子节点
				parentNode.ChangeToBlackNode()
				ppNode.ChangeToRedNode()
				node.ChangeToRedNode()
				if ppNode.IsRoot() {
					tr.Root = parentNode
				}
				ppNode.RightHand()
			} else {
				parentNode.RightHand()
				node.ChangeToBlackNode()
				ppNode.ChangeToRedNode()
				if ppNode.IsRoot() {
					tr.Root = parentNode
				}
				ppNode.LeftHand()
			}
		} else {
			node.ChangeToRedNode()
			parentNode.ChangeToBlackNode()
			uncleNode.ChangeToBlackNode()
			parentNode.ParentNode.ChangeToRedNode()
			if parentNode.ParentNode.ParentNode == nil {
				parentNode.ParentNode.ChangeToBlackNode()
				return nil
			}
			return tr.insertNode(parentNode.ParentNode, parentNode.ParentNode.ParentNode)
		}
	} else if node.Value > parentNode.Value {
		node.ParentNode = parentNode
		parentNode.RChildNode = node
		if !parentNode.IsRedNode() {
			node.ChangeToRedNode()
			node.ParentNode = parentNode
			parentNode.RChildNode = node
			return nil
		}
		uncleNode := parentNode.FindBrotherNode()
		if uncleNode == nil || !uncleNode.IsRedNode() {
			parentNode.ChangeToBlackNode()
			ppNode := parentNode.ParentNode
			if parentNode == ppNode.LChildNode {
				parentNode.LeftHand()
				node.ChangeToBlackNode()
				ppNode.ChangeToRedNode()
				if ppNode.IsRoot() {
					tr.Root = parentNode
				}
				ppNode.RightHand()
			} else {
				parentNode.ChangeToBlackNode()
				ppNode.ChangeToRedNode()
				node.ChangeToRedNode()
				if ppNode.IsRoot() {
					tr.Root = parentNode
				}
				ppNode.LeftHand()
			}
		} else {
			node.ChangeToRedNode()
			parentNode.ChangeToBlackNode()
			uncleNode.ChangeToBlackNode()
			parentNode.ParentNode.ChangeToRedNode()
			if parentNode.ParentNode.ParentNode == nil {
				parentNode.ParentNode.ChangeToBlackNode()
				return nil
			}
			return tr.insertNode(parentNode.ParentNode, parentNode.ParentNode.ParentNode)
		}
	}
	return nil
}

func (tr *BRTree) DeleteValue(value int) error {
	deleteNode, err := tr.FindNode(value)
	if err != nil {
		return err
	}
	parentNode := deleteNode.ParentNode
	if deleteNode.IsLeaf() {
		fmt.Printf("deleteNode is leaf %s \n", deleteNode.ToString())
		return tr.balanceWithDeleteNode(deleteNode)
	} else {
		if deleteNode.LChildNode != nil && deleteNode.RChildNode == nil {
			deleteNode.LChildNode.ParentNode = parentNode
			if parentNode.LChildNode == deleteNode {
				parentNode.LChildNode = deleteNode.LChildNode
				deleteNode.LChildNode.ChangeToDistNodeColor(deleteNode)
			} else {
				parentNode.RChildNode = deleteNode.LChildNode
				deleteNode.LChildNode.ChangeToDistNodeColor(deleteNode)
			}
		} else if deleteNode.LChildNode == nil && deleteNode.RChildNode != nil {
			deleteNode.RChildNode.ParentNode = parentNode
			if parentNode.LChildNode == deleteNode {
				parentNode.LChildNode = deleteNode.RChildNode
				deleteNode.RChildNode.ChangeToDistNodeColor(deleteNode)
			} else {
				parentNode.RChildNode = deleteNode.RChildNode
				deleteNode.RChildNode.ChangeToDistNodeColor(deleteNode)
			}
		} else {
			replaceNode := deleteNode.FindRChildMinNodeNode()
			deleteNode.Value, replaceNode.Value = replaceNode.Value, deleteNode.Value
			return tr.balanceWithDeleteNode(replaceNode)
		}
	}
	return nil
}

func (tr *BRTree) balanceWithDeleteNode(deleteNode *BRTreeNode) error {
	parentNode := deleteNode.ParentNode
	if parentNode == nil {
		return nil
	}
	if parentNode.LChildNode == deleteNode {
		if !deleteNode.IsRedNode() {
			brother := parentNode.RChildNode
			if brother.IsRedNode() {
				parentNode.ChangeToRedNode()
				brother.ChangeToBlackNode()
				parentNode.LeftHand()
				brother = parentNode.RChildNode
				brother.ChangeToRedNode()
				parentNode.LChildNode = nil
				parentNode.Value, deleteNode.Value = deleteNode.Value, parentNode.Value
				return tr.DeleteValue(parentNode.Value)
			} else {
				if brother.RChildIsRed() {
					brother.Color = parentNode.Color
					parentNode.ChangeToBlackNode()
					brother.RChildNode.ChangeToBlackNode()
					parentNode.LeftHand()
					parentNode.LChildNode = nil
				} else if !brother.RChildIsRed() && brother.LChildIsRed() {
					brother.ChangeToRedNode()
					brother.LChildNode.ChangeToBlackNode()
					brother.RightHand()
					brother = brother.ParentNode

					brother.Color = parentNode.Color
					parentNode.ChangeToBlackNode()
					brother.RChildNode.ChangeToBlackNode()
					parentNode.LeftHand()
					parentNode.LChildNode = nil
				} else if !brother.RChildIsRed() && !brother.LChildIsRed() {
					brother.ChangeToRedNode()
					parentNode.Value, deleteNode.Value = deleteNode.Value, parentNode.Value
					// 递归处理父节点
					return tr.DeleteValue(parentNode.Value)
				}
			}
		} else {
			parentNode.LChildNode = nil
			return nil
		}
	} else if parentNode.RChildNode == deleteNode {
		if !deleteNode.IsRedNode() {
			brother := parentNode.LChildNode
			if brother.IsRedNode() {
				brother.ChangeToBlackNode()
				parentNode.ChangeToRedNode()
				parentNode.RightHand()
				brother = parentNode.LChildNode
				brother.ChangeToRedNode()
				parentNode.Value, deleteNode.Value = deleteNode.Value, parentNode.Value
				return tr.DeleteValue(parentNode.Value)
			} else {
				if brother.LChildIsRed() {
					brother.Color = parentNode.Color
					brother.LChildNode.ChangeToBlackNode()
					parentNode.ChangeToBlackNode()
					parentNode.RightHand()
					parentNode.RChildNode = nil
				} else if !brother.LChildIsRed() && brother.RChildIsRed() {
					brother.ChangeToRedNode()
					brother.RChildNode.ChangeToBlackNode()
					brother.LeftHand()
					brother = brother.ParentNode

					brother.Color = parentNode.Color
					brother.LChildNode.ChangeToBlackNode()
					parentNode.ChangeToBlackNode()
					parentNode.RightHand()
					parentNode.RChildNode = nil
				} else if brother.LChildIsRed() && !brother.RChildIsRed() {
					brother.ChangeToRedNode()
					parentNode.Value, deleteNode.Value = deleteNode.Value, parentNode.Value
					return tr.DeleteValue(parentNode.Value)
				}
			}
		} else {
			parentNode.RChildNode = nil
		}
	}
	return nil
}
