package tree

import "fmt"

type Color string

type BRTreeNode struct {
	Value      int
	Color      Color
	LChildNode *BRTreeNode // left child node
	RChildNode *BRTreeNode // right child node
	ParentNode *BRTreeNode // parent node
}

const (
	Red   Color = "red"
	Black Color = "black"
)

func (node *BRTreeNode) IsRedNode() bool {
	return node.Color == Red
}

func (node *BRTreeNode) ToString() string {
	result := fmt.Sprintf("%s %d ", node.Color, node.Value)
	if node.LChildNode != nil {
		result = result + fmt.Sprintf("[left child %s %d] ", node.LChildNode.Color, node.LChildNode.Value)
	}
	if node.RChildNode != nil {
		result = result + fmt.Sprintf("[right child %s %d] ", node.RChildNode.Color, node.RChildNode.Value)
	}
	return result
}

func (node *BRTreeNode) ChangeToRedNode() {
	node.Color = Red
}

func (node *BRTreeNode) ChangeToBlackNode() {
	node.Color = Black
}

func (node *BRTreeNode) IsRoot() bool {
	return node.ParentNode == nil
}

func (node *BRTreeNode) FindBrotherNode() *BRTreeNode {
	parentNode := node.ParentNode
	if parentNode.LChildNode == node {
		return parentNode.RChildNode
	}
	return parentNode.LChildNode
}

func (node *BRTreeNode) LeftHand() {
	pRightChild := node.RChildNode
	pParent := node.ParentNode
	if pRightChild != nil {
		node.RChildNode = pRightChild.LChildNode
		if pRightChild.LChildNode != nil {
			pRightChild.LChildNode.ParentNode = node
		}
	}
	node.ParentNode = pRightChild
	pRightChild.LChildNode = node
	// 旋转节点不是根节点, 存在父节点
	if pParent != nil {
		if pParent.LChildNode == node {
			pParent.LChildNode = pRightChild
		} else {
			pParent.RChildNode = pRightChild
		}
	}
	pRightChild.ParentNode = pParent
}

func (node *BRTreeNode) RightHand() {
	pLeftChild := node.LChildNode
	pParent := node.ParentNode
	if pLeftChild != nil {
		node.LChildNode = pLeftChild.RChildNode
		if pLeftChild.RChildNode != nil {
			pLeftChild.RChildNode.ParentNode = node
		}
	}
	node.ParentNode = pLeftChild
	pLeftChild.RChildNode = node
	// 旋转节点不是根节点, 存在父节点
	if pParent != nil {
		if pParent.LChildNode == node {
			pParent.LChildNode = pLeftChild
		} else {
			pParent.RChildNode = pLeftChild
		}
	}
	pLeftChild.ParentNode = pParent
}
