package main

import (
	"fmt"

	tm "github.com/jfinken/trapezoidal-map"
)

func main() {

	root := &tm.Node{}
	pt := tm.Point{0, 0}
	fmt.Printf("%v\n", root.Search(&pt))

	// UL to LR: X increasing, Y increasing (down)

	// TODO: generate test sets of segments
	s1 := &tm.Segment{P: &tm.Point{66, 192}, Q: &tm.Point{318, 152}}
	s2 := &tm.Segment{P: &tm.Point{200, 256}, Q: &tm.Point{437, 282}}
	segs := []*tm.Segment{s1, s2}
	tm.ConstructMap(1024, 1024, segs)
}
