package trapezoidalmap

import (
	"testing"
)

func TestAbove(t *testing.T) {

	seg := &Segment{
		P: &Point{X: 10, Y: 10},
		Q: &Point{X: 30, Y: 15}}
	pt := &Point{X: 25, Y: 27}

	got := seg.Above(pt)

	if got != true {
		t.Errorf("trapezoidalmap.Above error! Got %t, Expected: %t", got, true)
	}

	pt = &Point{X: 25, Y: 7}
	got = seg.Above(pt)

	if got == true {
		t.Errorf("trapezoidalmap.Above error! Got %t, Expected: %t", got, false)
	}
}

func TestRight(t *testing.T) {
	p1 := &Point{X: 25, Y: 127}
	p2 := &Point{X: 50, Y: 270}

	// is p1 to the right of p2
	got := p1.Right(p2)

	if got == true {
		t.Errorf("trapezoidalmap.Right error! Got %t, Expected: %t", got, false)
	}

	got = p2.Right(p1)

	if got != true {
		t.Errorf("trapezoidalmap.Right error! Got %t, Expected: %t", got, true)
	}
}
