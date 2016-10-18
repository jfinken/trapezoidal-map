package trapezoidalmap

type Type int

const (
	Leaf  Type = iota // 0
	XNode             // 1
	YNode             // 2
)

type Node struct {
	Type   Type
	Parent *Node
	Left   *Node
	Right  *Node
	P      *Point
	S      *Segment
	T      *Trapezoid
}

// Search the tree starting at n (typically the root) for p and
// return a pointer to the Trapezoid containing p or nil if none exists.
func (n *Node) Search(p *Point) *Trapezoid {

	next := n
	for next.Type != Leaf && next != nil {
		next = next.traverse(p)
	}
	return next.T
}
func (n *Node) traverse(p *Point) *Node {

	if n.Type == XNode { // x-node
		if p.Right(n.P) {
			return n.Right
		}
		return n.Left
	} else if n.Type == YNode { // y-node
		if n.S.Above(p) {
			return n.Left
		}
		return n.Right
	}
	return nil
}
func (n *Node) setRightTree(rightPt *Point, rightTrap *Trapezoid, leftSeg *Segment,
	leftSubTrapLeft *Trapezoid, leftSubTrapRight *Trapezoid) {

	// Ref: Figure 6.7, page 129 of http://www.cs.uu.nl/geobook/
	n.Right = &Node{P: rightPt, Parent: n, Type: XNode}
	n.Right.Right = &Node{T: rightTrap, Parent: n.Right, Type: Leaf}

	n.Right.Left = &Node{S: leftSeg, Parent: n.Right, Type: YNode}
	n.Right.Left.Left = &Node{T: leftSubTrapLeft, Parent: n.Right.Left, Type: Leaf}
	n.Right.Left.Right = &Node{T: leftSubTrapRight, Parent: n.Right.Left, Type: Leaf}
}