package trapezoidalmap

import (
	"math/rand"
	"time"
)

// ConstructMap is the main entry point to construct a trapezoidal map
// and its interlinked search data structure D.  Width and Height define
// the over-arching bounding box R.
func ConstructMap(width, height int, segments []*Segment) {

	// start T and D with a nil trapezoid (R?)
	t0 := &Trapezoid{}
	trapMap := []*Trapezoid{t0}
	// interlinking
	root := &Node{T: t0, Type: Leaf}
	t0.Node = root

	for len(segments) > 0 {
		r := random(len(segments))
		seg := segments[r]
		segments = removeSeg(segments, r)

		// find the set of trapezoids in T intersected by seg
		intersectedTraps := followSegment(root, seg)

		// simple case: the segment is completely contained in trapezoid 0
		if len(intersectedTraps) == 1 {

			d := intersectedTraps[0]
			trapMap = removeTrap(trapMap, 0)

			// set end segments and points
			A := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: d.Leftp, Rightp: seg.P}
			C := &Trapezoid{Top: d.Top, Bottom: seg, Leftp: seg.P, Rightp: seg.Q}
			D := &Trapezoid{Top: seg, Bottom: d.Bottom, Leftp: seg.P, Rightp: seg.Q}
			B := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: seg.Q, Rightp: d.Rightp}

			// set neighbors: UL, LL, UR, LR
			A.setNeighbors(d.UpperLeft, d.LowerLeft, C, D)
			C.setNeighbors(A, A, B, B)
			D.setNeighbors(A, A, B, B)
			B.setNeighbors(C, D, d.UpperRight, d.LowerRight)

			if d.UpperLeft != nil {
				d.UpperLeft.UpperRight = A
				d.UpperLeft.LowerRight = A
			}
			if d.LowerLeft != nil {
				d.LowerLeft.UpperRight = A
				d.LowerLeft.LowerRight = A
			}
			if d.UpperRight != nil {
				d.UpperRight.UpperLeft = B
				d.UpperRight.LowerLeft = B
			}
			if d.LowerRight != nil {
				d.LowerRight.UpperLeft = B
				d.LowerRight.LowerLeft = B
			}
			// update the tree D by replacing the leaf for d by a little
			// tree with four leaves.
			subTree := d.Node
			subTree.P = seg.P
			subTree.Type = XNode

			// set left sub-tree and interlink it with the trapezoid
			subTree.Left = &Node{T: A, Type: Leaf}
			subTree.Left.Parent = subTree
			A.Node = subTree
			// set right sub-tree
			subTree.setRightTree(seg.Q, B, seg, C, D)
		} else {
			// newTrapezoids :=
			// prevUpper :=
			// prevLower :=

			// much more complicated case: seg intersects two or more trapezoids.
			for i := 0; i < len(intersectedTraps); i++ {
				d := intersectedTraps[i]

				if i == 0 {
					trapMap = removeTrap(trapMap, i)
					A := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: d.Leftp, Rightp: seg.P}
					B := &Trapezoid{Top: d.Top, Bottom: seg, Leftp: seg.P, Rightp: d.Rightp}
					C := &Trapezoid{Top: seg, Bottom: d.Bottom, Leftp: seg.P, Rightp: d.Rightp}
					A.setNeighbors(d.UpperLeft, d.LowerLeft, B, C)
					B.setNeighbors(A, A, nil, nil)
					C.setNeighbors(A, A, nil, nil)
				}
			}
		}
	}
}
func removeSeg(segs []*Segment, index int) []*Segment {

	copy(segs[index:], segs[index+1:])
	// This is a slice of pointers, allow GC
	segs[len(segs)-1] = nil
	return segs[:len(segs)-1]
}

// correct: no generics
func removeTrap(traps []*Trapezoid, index int) []*Trapezoid {

	copy(traps[index:], traps[index+1:])
	// This is a slice of pointers, this allows GC
	traps[len(traps)-1] = nil
	return traps[:len(traps)-1]
}

// followSegment searches for and returns the slice of trapezoids intersected
// by the Segment seg.
func followSegment(root *Node, seg *Segment) []*Trapezoid {

	var traversed []*Trapezoid
	p := seg.P
	q := seg.Q

	// First search the graph for p
	// TBD: how/where should the root of the search data structure be stored?
	t0 := root.Search(p)
	if t0 == nil {
		return traversed
	}

	traversed = append(traversed, t0)
	j := t0

	for j != nil && (j.Rightp != nil && q.Right(j.Rightp)) {
		if seg.Above(j.Rightp) {
			j = j.LowerRight
		} else {
			j = j.UpperRight
		}
		if j != nil {
			traversed = append(traversed, j)
		}
	}
	return traversed
}

// TrapezoidMap is a randomized incremental algorithm.  random here
// returns an int to be used to index into the slice of Segments s.
func random(max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return rnd.Intn(max)
}
