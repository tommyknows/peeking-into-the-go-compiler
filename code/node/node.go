package syntax

import "go/types"

// A Node is a single node in the syntax tree.
type Node struct {
	// Tree structure.
	// Generic recursive walks should follow these fields.
	Left  *Node
	Right *Node
	Ninit Nodes
	Nbody Nodes
	List  Nodes
	Rlist Nodes

	// most nodes
	Type *types.Type

	Pos src.XPos
	// ...and more
}
