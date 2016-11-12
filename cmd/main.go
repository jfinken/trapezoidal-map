package main

import tm "github.com/jfinken/trapezoidal-map"

func main() {

	//root := &tm.Node{}
	//pt := tm.Point{0, 0}
	//fmt.Printf("%v\n", root.Search(&pt))

	// UL to LR: X increasing, Y increasing (down)

	// TODO: generate test sets of segments

	/* SIMPLE
	s1 := &tm.Segment{P: &tm.Point{66, 192}, Q: &tm.Point{318, 152}, Index: 0}
	s2 := &tm.Segment{P: &tm.Point{200, 256}, Q: &tm.Point{437, 282}, Index: 1}
	segs := []*tm.Segment{s1, s2}
	*/

	// complex
	s1 := &tm.Segment{P: &tm.Point{87, 159}, Q: &tm.Point{335, 100}, Index: 0}
	s2 := &tm.Segment{P: &tm.Point{215, 222}, Q: &tm.Point{422, 259}, Index: 1}
	s3 := &tm.Segment{P: &tm.Point{19, 332}, Q: &tm.Point{133, 325}, Index: 2}
	s4 := &tm.Segment{P: &tm.Point{119, 245}, Q: &tm.Point{256, 406}, Index: 3}
	segs := []*tm.Segment{s1, s2, s3, s4}

	trapMap := tm.ConstructMap(1024, 1024, segs)

	//log.Printf("RESULT: %v\n", trapMap)

	tm.RenderProcessing(trapMap, 512, 512)
}
