package trapezoidalmap

type Point struct {
	X float64
	Y float64
}
type Segment struct {
	P *Point
	Q *Point
}
type Trapezoid struct {
	Top    *Segment
	Bottom *Segment
	Leftp  *Point
	Rightp *Point
	Node   *Node

	UpperLeft  *Trapezoid
	LowerLeft  *Trapezoid
	UpperRight *Trapezoid
	LowerRight *Trapezoid
}

// setNeighbors is a convenience function
func (t *Trapezoid) setNeighbors(ul, ll, ur, lr *Trapezoid) {
	t.UpperLeft = ul
	t.LowerLeft = ll
	t.UpperRight = ur
	t.LowerRight = lr
}

// Right is an X-Node query.  It returns whether or not the given point is to
// the right of p.
func (p *Point) Right(that *Point) bool {
	return (p.X >= that.X)
}

// Above is a Y-Node query.  It returns whether or not p lies above or below
// the Segment s.
func (s *Segment) Above(p *Point) bool {
	return (p.Y > line(s.P, s.Q, p))
}

func line(left, right, pt *Point) float64 {
	return (((right.Y - left.Y) / (right.X - left.X)) * (pt.X - left.X)) + left.Y
}
