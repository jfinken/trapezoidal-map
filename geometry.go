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

func (t *Trapezoid) equals(other *Trapezoid) bool {
	if other.isNil() && t.isNil() {
		return true
	}
	// everything else: Segments, Points, Neighbors
	return (t.Top != nil && t.Top.equals(other.Top)) &&
		(t.Bottom != nil && t.Bottom.equals(other.Bottom)) &&
		(t.Leftp != nil && t.Leftp.equals(other.Leftp)) &&
		(t.Rightp != nil && t.Rightp.equals(other.Rightp)) &&
		(t.UpperLeft.equals(other.UpperLeft)) &&
		(t.UpperRight.equals(other.UpperRight)) &&
		(t.LowerLeft.equals(other.LowerLeft)) &&
		(t.LowerRight.equals(other.LowerRight))
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
func (p *Point) equals(other *Point) bool {
	return (p.X == other.X && p.Y == other.Y)
}
func (s *Segment) equals(other *Segment) bool {
	return (s.P == other.P && s.Q == other.Q)
}
func (t *Trapezoid) isNil() bool {
	return (t.Top == nil && t.Bottom == nil && t.Leftp == nil && t.Rightp == nil)
}
