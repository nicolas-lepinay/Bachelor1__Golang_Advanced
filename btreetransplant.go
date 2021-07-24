package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	// Si la racine est inexistante, il n'y a pas d'arbre, donc on renvoit nil :
	if root == nil {
		return nil
	}

	// Si la node de remplacement (rplc) est supérieure au parent de la node à remplacer (node)...
	if rplc.Data > node.Parent.Data {
		//.. alors rplc prend la place de l'enfant droite du parent de la node à remplacer.
		node.Parent.Right = rplc
	}
	// Si la node de remplacement (rplc) est inférieure au parent de la node à remplacer (node)...
	if rplc.Data < node.Parent.Data {
		//.. alors rplc prend la place de l'enfant gauche du parent de la node à remplacer.
		node.Parent.Left = rplc
	}

	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root

		// If data > root.Data :
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data || root.Data == "" {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
