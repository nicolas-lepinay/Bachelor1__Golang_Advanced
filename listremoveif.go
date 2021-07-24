package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {

	// Si la liste est vide, ou si la liste ne contient qu'un seul node dont la donnée est data_ref, on affiche "nil" :
	if l == nil || l.Tail == l.Head && l.Head.Data == data_ref {
		l.Head = nil
		return
	}

	// On vérifie si le 1er node contient data_ref. Si c'est le cas, on le supprime en passant au node suivant ("l.Head = l.Head.Next") :
	for l.Head.Data == data_ref {
		if l.Head.Data == data_ref {
			l.Head = l.Head.Next
		}
	}

	// On supprime data_ref de la liste grâce à cette partie :
	node := l.Head
	for node.Next != nil {

		// Si la donnée d'un node est égale à data_ref, alors ce node prend pour valeur le node suivant :
		for node.Next.Data == data_ref {
			if node.Next.Data == data_ref {
				node.Next = node.Next.Next
			}
		}
		// Si la boucle 'for' n'est pas exécutée une première fois (c-à-d si node.Next.Data n'est pas égale à data_ref), ou bien après qu'elle se soit exécutée, on passe au node suivant en faisant node = node.Next :
		node = node.Next

		// Si on est en train de parcourir le dernier node et qu'il contient data_ref, on le remplace par "nil" :
		if node.Next == l.Tail && node.Next.Data == data_ref {
			node.Next = nil
			return
		}
	}

}

func ListPushBack(l *List, data interface{}) {

	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode

	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
