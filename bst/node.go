package bst

// Node представляет узел бинарного дерева поиска
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
}