package tree

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

func (node *BRTreeNode) ChangeToRedNode() {
	node.Color = Red
}

func (node *BRTreeNode) ChangeToBlackNode() {
	node.Color = Black
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
	node.RChildNode = pRightChild.LChildNode
	pRightChild.LChildNode.ParentNode = node
	node.ParentNode = pRightChild
	pRightChild.LChildNode = node
	if pParent.LChildNode == node {
		pParent.LChildNode = pRightChild
	} else {
		pParent.RChildNode = pRightChild
	}
	pRightChild.ParentNode = pParent
}

func (node *BRTreeNode) RightHand() {
	pLeftChild := node.LChildNode
	pParent := node.ParentNode
	node.LChildNode = pLeftChild.RChildNode
	pLeftChild.RChildNode.ParentNode = node
	node.ParentNode = pLeftChild
	pLeftChild.RChildNode = node
	if pParent.LChildNode == node {
		pParent.LChildNode = pLeftChild
	} else {
		pParent.RChildNode = pLeftChild
	}
	pLeftChild.ParentNode = pParent
}
