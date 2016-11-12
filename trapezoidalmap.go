package trapezoidalmap

import (
	"fmt"
	"math/rand"
	"time"
)

// ConstructMap is the main entry point to construct a trapezoidal map
// and its interlinked search data structure D.  Width and Height define
// the over-arching bounding box R.
func ConstructMap(width, height int, segments []*Segment) []*Trapezoid {

	// start T and D with a nil trapezoid (R?)
	t0 := &Trapezoid{}
	// interlinking
	root := &Node{T: t0, Type: Leaf}
	t0.Node = root
	trapMap := []*Trapezoid{t0}

	for len(segments) > 0 {
		r := random(len(segments))
		seg := segments[r]
		segments = removeSeg(segments, r)

		// find the set of trapezoids in T intersected by seg
		intersectedTraps := followSegment(root, seg)

		fmt.Printf("Segment: %d\n", seg.Index)
		fmt.Printf("Intersected: %d\n", len(intersectedTraps))

		// simple case: the segment is completely contained in trapezoid 0
		if len(intersectedTraps) == 1 {

			d := intersectedTraps[0]
			trapMap = removeTrap(trapMap, d)

			// set end segments and points
			A := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: d.Leftp, Rightp: seg.P, UUID: newUUID()}
			C := &Trapezoid{Top: d.Top, Bottom: seg, Leftp: seg.P, Rightp: seg.Q, UUID: newUUID()}
			D := &Trapezoid{Top: seg, Bottom: d.Bottom, Leftp: seg.P, Rightp: seg.Q, UUID: newUUID()}
			B := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: seg.Q, Rightp: d.Rightp, UUID: newUUID()}

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
			// add to trapezoidal map
			trapMap = append(trapMap, A, B, C, D)

			// update the tree D by replacing the leaf for d by a little
			// tree with four leaves.
			d.Node.P = seg.P
			d.Node.Type = XNode

			// set left sub-tree and interlink it with the trapezoid
			d.Node.setLeftTree(A)

			// set right sub-tree
			d.Node.setRightTreeToPoint(seg.Q, B, seg, C, D)

		} else {
			newTrapezoids := []*Trapezoid{}
			var prevUpper *Trapezoid
			var prevLower *Trapezoid

			// much more complicated case: seg intersects two or more trapezoids.
			for i := 0; i < len(intersectedTraps); i++ {
				d := intersectedTraps[i]
				fmt.Printf("d: %v\n", d)
				trapMap = removeTrap(trapMap, d)

				if i == 0 {
					A := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: d.Leftp, Rightp: seg.P, UUID: newUUID()}
					B := &Trapezoid{Top: d.Top, Bottom: seg, Leftp: seg.P, Rightp: d.Rightp, UUID: newUUID()}
					C := &Trapezoid{Top: seg, Bottom: d.Bottom, Leftp: seg.P, Rightp: d.Rightp, UUID: newUUID()}
					A.setNeighbors(d.UpperLeft, d.LowerLeft, B, C)
					B.setNeighbors(A, A, nil, nil)
					C.setNeighbors(A, A, nil, nil)

					if d.UpperLeft != nil {
						d.UpperLeft.UpperRight = A
						d.UpperLeft.LowerRight = A
					}
					if d.LowerLeft != nil {
						d.LowerLeft.LowerRight = A
						d.LowerLeft.UpperRight = A
					}

					trapMap = append(trapMap, A)
					newTrapezoids = append(newTrapezoids, B, C)
					prevUpper = B
					prevLower = C

					// update the tree D by replacing the leaf for d by a little
					// tree with four leaves.
					d.Node.P = seg.P
					d.Node.Type = XNode

					// set left sub-tree and interlink it with the trapezoid
					d.Node.setLeftTree(A)

					// set right sub-tree
					d.Node.setRightTreeToSeg(seg, B, C)

					// for each trapezoid in the map, check trap's neighbors
					assertNeighbors(trapMap)

				} else if i == len(intersectedTraps)-1 {
					B := &Trapezoid{Top: d.Top, Bottom: seg, Leftp: d.Leftp, Rightp: seg.Q, UUID: newUUID()}
					C := &Trapezoid{Top: seg, Bottom: d.Bottom, Leftp: d.Leftp, Rightp: seg.Q, UUID: newUUID()}
					A := &Trapezoid{Top: d.Top, Bottom: d.Bottom, Leftp: seg.Q, Rightp: d.Rightp, UUID: newUUID()}
					B.setNeighbors(prevUpper, prevUpper, A, A)
					C.setNeighbors(prevLower, prevLower, A, A)
					A.setNeighbors(B, C, d.UpperRight, d.LowerRight)

					prevUpper.UpperRight = B
					prevUpper.LowerRight = B
					prevLower.UpperRight = C
					prevLower.LowerRight = C

					if d.UpperRight != nil {
						d.UpperRight.UpperLeft = A
						d.UpperRight.LowerLeft = A
					}
					if d.LowerRight != nil {
						d.LowerRight.LowerLeft = A
						d.LowerRight.UpperLeft = A
					}
					trapMap = append(trapMap, A)
					newTrapezoids = append(newTrapezoids, B, C)

					// update the tree D by replacing the leaf for d by a little
					// tree with four leaves.
					d.Node.P = seg.Q
					d.Node.Type = XNode

					// set right sub-tree and interlink it with the trapezoid
					d.Node.setRightTree(A)

					// set left sub-tree
					d.Node.setLeftTreeToSeg(seg, B, C)

					// for each trapezoid in the map, check trap's neighbors
					assertNeighbors(trapMap)
				} else {

					A := &Trapezoid{Top: d.Top, Bottom: seg, Leftp: d.Leftp, Rightp: d.Rightp, UUID: newUUID()}
					B := &Trapezoid{Top: seg, Bottom: d.Bottom, Leftp: d.Leftp, Rightp: d.Rightp, UUID: newUUID()}
					A.setNeighbors(prevUpper, prevUpper, nil, nil)
					B.setNeighbors(prevLower, prevLower, nil, nil)

					prevUpper.UpperRight = A
					prevUpper.LowerRight = A
					prevLower.UpperRight = B
					prevLower.LowerRight = B

					prevUpper = A
					prevLower = B
					newTrapezoids = append(newTrapezoids, A, B)
					// update the tree D by replacing the leaf for d by a little
					// tree with four leaves.
					d.Node.S = seg
					d.Node.Type = YNode

					// set left sub-tree and interlink it with the trapezoid
					d.Node.setLeftTree(A)

					// set right sub-tree
					d.Node.setRightTree(B)

					// for each trapezoid in the map, check trap's neighbors
					assertNeighbors(trapMap)
				}
			}
			newTrapezoids = mergeTrapezoids(newTrapezoids, seg)
			// reset all
			for _, t := range newTrapezoids {
				t.Merged = false
				trapMap = append(trapMap, t)
			}
		}
	}
	return trapMap
}
func removeSeg(segs []*Segment, index int) []*Segment {

	copy(segs[index:], segs[index+1:])
	// This is a slice of pointers, allow GC
	segs[len(segs)-1] = nil
	return segs[:len(segs)-1]
}

// correct: no generics
func removeTrap(traps []*Trapezoid, t *Trapezoid) []*Trapezoid {

	// find the item to remove
	for i := range traps {

		if traps[i].equals(t) {
			copy(traps[i:], traps[i+1:])
			// This is a slice of pointers, this allows GC
			traps[len(traps)-1] = nil
			return traps[:len(traps)-1]
		}
	}
	fmt.Printf("Removing NO traps!\n")
	return traps
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

func mergeTrapezoids(newTrapezoids []*Trapezoid, seg *Segment) []*Trapezoid {
	merged := false
	for !merged {
		for i, t := range newTrapezoids {
			if t.Rightp != nil && !t.Rightp.equals(seg.P) && !t.Rightp.equals(seg.Q) &&
				((t.Top != nil && t.Top.Above(t.Rightp)) ||
					(t.Bottom != nil && !t.Bottom.Above(t.Rightp))) {

				next := t.UpperRight
				nextI := i
				t.UpperRight = next.UpperRight
				t.LowerRight = next.LowerRight

				if t.Top != nil && t.Top.Above(t.Rightp) {
					t.UpperRight.LowerLeft = t
				} else {
					t.UpperRight.UpperLeft = t
				}
				t.Rightp = next.Rightp
				// update the search structure
				if next.Node.Parent.isLeftNode(next.Node) {
					next.Node.Parent.Left = t.Node
				} else {
					next.Node.Parent.Right = t.Node
				}
				newTrapezoids = removeTrap(newTrapezoids, newTrapezoids[nextI])
				break
			} else {
				t.Merged = true
			}
		}
		//?
		merged = true
		for _, t := range newTrapezoids {
			if !t.Merged {
				merged = false
			}
		}
	}
	return newTrapezoids
}

// TrapezoidMap is a randomized incremental algorithm.  random here
// returns an int to be used to index into the slice of Segments s.
func random(max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return rnd.Intn(max)
}
func assertNeighbors(traps []*Trapezoid) {
	for _, t := range traps {
		if contains(traps, t.UpperLeft) == false {
			t.UpperLeft = nil
		}
		if contains(traps, t.LowerLeft) == false {
			t.LowerLeft = nil
		}
		if contains(traps, t.UpperRight) == false {
			t.UpperRight = nil
		}
		if contains(traps, t.LowerRight) == false {
			t.LowerRight = nil
		}
	}
}
func contains(traps []*Trapezoid, trap *Trapezoid) bool {
	for _, t := range traps {
		if t == trap {
			return true
		}
	}
	return false
}
