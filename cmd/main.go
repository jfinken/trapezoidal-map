package main

import (
	"fmt"

	tm "github.com/jfinken/trapezoidal-map"
)

func main() {

	root := &tm.Node{}
	pt := tm.Point{0, 0}
	fmt.Printf("%v\n", root.Search(&pt))

	// TODO: generate test sets of segments
	s1 := &tm.Segment{P: &tm.Point{128, 256}, Q: &tm.Point{256, 300}}
	_ = &tm.Segment{P: &tm.Point{200, 512}, Q: &tm.Point{400, 612}}
	segs := []*tm.Segment{s1}
	tm.ConstructMap(1024, 1024, segs)
}
