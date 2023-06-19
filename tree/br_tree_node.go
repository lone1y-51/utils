package tree

import (
	"fmt"

	"github.com/lone1y-51/utils/define"
)

type Color string

type BRTreeNode[T define.Sorted] struct {
	Value      T
	Color      Color
	LChildNode *BRTreeNode[T] // left child node
	RChildNode *BRTreeNode[T] // right child node
	ParentNode *BRTreeNode[T] // parent node
	// LChildNode *BRTreeNode // left child node
	// RChildNode *BRTreeNode // right child node
	// ParentNode *BRTreeNode // parent node
}

const (
	Red   Color = "red"
	Black Color = "black"
)

func (node *BRTreeNode[T]) IsRoot() bool {
	return node.ParentNode == nil
}

func (node *BRTreeNode[T]) IsLeaf() bool {
	if node.LChildNode == nil && node.RChildNode == nil {
		return true
	}
	return false
}

func (node *BRTreeNode[T]) IsRedNode() bool {
	return node.Color == Red
}

// 子节点不存在认为是黑色
func (node *BRTreeNode[T]) LChildIsRed() bool {
	if node.LChildNode != nil && node.LChildNode.IsRedNode() {
		return true
	}
	return false
}

func (node *BRTreeNode[T]) RChildIsRed() bool {
	if node.RChildNode != nil && node.RChildNode.IsRedNode() {
		return true
	}
	return false
}

func (node *BRTreeNode[T]) ChangeToDistNodeColor(dst *BRTreeNode[T]) {
	node.Color = dst.Color
}

func (node *BRTreeNode[T]) ToString() string {
	result := fmt.Sprintf("%s %v ", node.Color, node.Value)
	if node.LChildNode != nil {
		result = result + fmt.Sprintf("[left child %s %v] ", node.LChildNode.Color, node.LChildNode.Value)
	}
	if node.RChildNode != nil {
		result = result + fmt.Sprintf("[right child %s %v] ", node.RChildNode.Color, node.RChildNode.Value)
	}
	return result
}

func (node *BRTreeNode[T]) ChangeToRedNode() {
	node.Color = Red
}

func (node *BRTreeNode[T]) ChangeToBlackNode() {
	node.Color = Black
}

func (node *BRTreeNode[T]) FindRChildMinNodeNode() *BRTreeNode[T] {
	result := node.RChildNode
	for {
		if result.LChildNode != nil {
			result = result.LChildNode
			continue
		}
		break
	}
	return result
}

func (node *BRTreeNode[T]) FindBrotherNode() *BRTreeNode[T] {
	parentNode := node.ParentNode
	if parentNode.LChildNode == node {
		return parentNode.RChildNode
	}
	return parentNode.LChildNode
}

func (node *BRTreeNode[T]) LeftHand() {
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

func (node *BRTreeNode[T]) RightHand() {
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
