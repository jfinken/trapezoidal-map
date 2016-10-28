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
func (n *Node) isLeftNode(other *Node) bool {
	if other == nil {
		return (n.P == nil && n.S == nil && n.T == nil)
	}
	if n.Left.equals(other) {
		return true
	}
	return false
}
func (n *Node) setRightTree(A *Trapezoid) {

	n.Right = &Node{T: A, Type: Leaf}
	n.Right.Parent = n
	A.Node = n
}
func (n *Node) setLeftTree(A *Trapezoid) {

	n.Left = &Node{T: A, Type: Leaf}
	n.Left.Parent = n
	A.Node = n
}
func (n *Node) setRightTreeToPoint(rightPt *Point, rightTrap *Trapezoid, leftSeg *Segment,
	leftSubTrapLeft *Trapezoid, leftSubTrapRight *Trapezoid) {

	// Ref: Figure 6.7, page 129 of http://www.cs.uu.nl/geobook/
	n.Right = &Node{P: rightPt, Parent: n, Type: XNode}
	n.Right.Right = &Node{T: rightTrap, Parent: n.Right, Type: Leaf}

	n.Right.Left = &Node{S: leftSeg, Parent: n.Right, Type: YNode}
	n.Right.Left.Left = &Node{T: leftSubTrapLeft, Parent: n.Right.Left, Type: Leaf}
	n.Right.Left.Right = &Node{T: leftSubTrapRight, Parent: n.Right.Left, Type: Leaf}
}
func (n *Node) setRightTreeToSeg(seg *Segment, leftTrap *Trapezoid, rightTrap *Trapezoid) {

	n.Right = &Node{S: seg, Parent: n, Type: YNode}
	n.Right.Left = &Node{T: leftTrap, Parent: n.Right, Type: Leaf}
	n.Right.Right = &Node{T: rightTrap, Parent: n.Right, Type: Leaf}
}

func (n *Node) setLeftTreeToSeg(seg *Segment, leftTrap *Trapezoid, rightTrap *Trapezoid) {

	n.Left = &Node{S: seg, Parent: n, Type: YNode}
	n.Left.Left = &Node{T: leftTrap, Parent: n.Right, Type: Leaf}
	n.Left.Right = &Node{T: rightTrap, Parent: n.Right, Type: Leaf}
}
func (n *Node) equals(other *Node) bool {

	if other == nil {
		return false
	}

	if other == nil && n.P == nil && n.S == nil && n.T == nil {
		return true
	}
	if n.P != nil && other.P != nil && !n.P.equals(other.P) {
		return false
	}
	if n.S != nil && other.S != nil && !n.S.equals(other.S) {
		return false
	}
	if n.T != nil && other.T != nil && !n.T.equals(other.T) {
		return false
	}
	if n.Type != other.Type {
		return false
	}
	return true
}
