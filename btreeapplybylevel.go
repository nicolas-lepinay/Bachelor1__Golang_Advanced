package student

type TreeNode struct {
	Parent, Left, Right *TreeNode
	Data                string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// depth : profondeur (nombre d'étages) de l'arbre
	depth := BTreeLevelCount(root)

	for i := 1; i <= depth; i++ {
		Level(root, i, f)
	}
}

func Level(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if level == 1 {
		f(root.Data)

	} else if level > 1 {
		Level(root.Left, level-1, f)
		Level(root.Right, level-1, f)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	countLeft := 0
	countRight := 0

	if root == nil {
		return 0

	} else {
		countLeft = BTreeLevelCount(root.Left)
		countLeft++

		countRight = BTreeLevelCount(root.Right)
		countRight++
	}

	if countLeft > countRight {
		return countLeft
	}
	return countRight
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root

	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
