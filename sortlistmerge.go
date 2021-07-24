package student

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {

	var tableau []int

	// On ajoute les données de chacune des 2 listes dans un tableau :
	for n1 != nil {
		tableau = append(tableau, n1.Data)
		n1 = n1.Next
	}

	for n2 != nil {
		tableau = append(tableau, n2.Data)
		n2 = n2.Next
	}

	// On trie le tableau par ordre croissant :
	for i := 0; i < len(tableau); i++ {
		for j := i + 1; j < len(tableau); j++ {
			if tableau[i] > tableau[j] {
				tableau[i], tableau[j] = tableau[j], tableau[i]
			}
		}
	}

	// On crée une liste à partir des valeurs du tableau :
	var nextNode *NodeI = nil
	for i := len(tableau) - 1; i >= 0; i-- {
		var node NodeI
		node.Data = tableau[i]
		node.Next = nextNode
		nextNode = &node
	}
	return nextNode
}
