package student

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {

	var tableau []int

	// On ajoute toutes les données de la liste dans un tableau :
	for l != nil {
		tableau = append(tableau, l.Data)
		l = l.Next
	}

	// A la fin du tableau, on rajoute la donnée à ajouter (data_ref) :
	tableau = append(tableau, data_ref)

	// On trie le tableau par ordre croissant :
	for i := 0; i < len(tableau); i++ {
		for j := i + 1; j < len(tableau); j++ {
			if tableau[i] > tableau[j] {
				tableau[i], tableau[j] = tableau[j], tableau[i]
			}
		}
	}

	var nextNode *NodeI = nil
	for i := len(tableau) - 1; i >= 0; i-- {
		var node NodeI
		node.Data = tableau[i]
		node.Next = nextNode
		nextNode = &node
	}

	return nextNode
}
