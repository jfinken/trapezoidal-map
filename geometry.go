package trapezoidalmap

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

type Point struct {
	X float64
	Y float64
}
type Segment struct {
	P     *Point
	Q     *Point
	Index int
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

	UUID   string
	Merged bool
}

func (t *Trapezoid) equals(other *Trapezoid) bool {
	if other == nil && t == nil {
		return true
	}
	if other.isNil() && t.isNil() {
		return true
	}
	if other.UUID == t.UUID {
		return true
	}

	return false
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
func (s *Segment) isNil() bool {
	return (s.P == nil && s.Q == nil)
}
func (s *Segment) equals(other *Segment) bool {
	return (s.P == other.P && s.Q == other.Q)
}
func (t *Trapezoid) isNil() bool {
	return (t.Top == nil && t.Bottom == nil && t.Leftp == nil && t.Rightp == nil)
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		// effectively swallow the error, not idiomatic
		log.Printf("%s\n", err.Error())
		return ""
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
