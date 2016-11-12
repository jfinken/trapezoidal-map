package trapezoidalmap

import (
	"bufio"
	"fmt"
	"os"
)

func print(str string) {
	fmt.Printf(str + "\n")
}
func printf(w *bufio.Writer, str string) {
	// TODO: error handling
	_, _ = w.WriteString(str + "\n")
	w.Flush()
}

// RenderProcessing writes the map geometry to a file suitable
// for pasting into your local copy of Processing
func RenderProcessing(tm []*Trapezoid, width, height int) {

	/*
	   type Trapezoid struct {
	   	Top    *Segment
	   	Bottom *Segment
	   	Leftp  *Point
	   	Rightp *Point
	   	...
	*/
	f, _ := os.Create("trapezoidalmap.processing.out")
	w := bufio.NewWriter(f)
	printf(w, fmt.Sprintf("void setup() {size(%d,%d); background(50);}", width, height))
	printf(w, fmt.Sprintf("void draw() {"))
	for _, t := range tm {

		// TOP
		if t.Top != nil {
			printf(w, fmt.Sprintf("//t.Top"))
			printf(w, fmt.Sprintf("stroke(255);"))
			printf(w, fmt.Sprintf("line(%f, %f, %f, %f);",
				t.Top.P.X, t.Top.P.Y, t.Top.Q.X, t.Top.Q.Y))
			printf(w, fmt.Sprintf("noStroke();"))
		}
		// Bottom
		if t.Bottom != nil {
			printf(w, fmt.Sprintf("//t.Bottom"))
			printf(w, fmt.Sprintf("stroke(255);"))
			printf(w, fmt.Sprintf("line(%f, %f, %f, %f);",
				t.Bottom.P.X, t.Bottom.P.Y, t.Bottom.Q.X, t.Bottom.Q.Y))
			printf(w, fmt.Sprintf("noStroke();"))
		}
		//Leftp
		if t.Leftp != nil {
			printf(w, fmt.Sprintf("//t.Leftp: green"))
			printf(w, fmt.Sprintf("fill(0, 255, 0);"))
			printf(w, fmt.Sprintf("ellipse(%f, %f, 5.0, 5.0);",
				t.Leftp.X, t.Leftp.Y))
			printf(w, fmt.Sprintf("noFill();"))

			// the vertical through this point
			printf(w, fmt.Sprintf("stroke(0, 255, 0);"))
			printf(w, fmt.Sprintf("line(%f, %f, %f, %f);",
				t.Leftp.X, t.Leftp.Y, t.Leftp.X, 0.0))
			printf(w, fmt.Sprintf("noStroke();"))
		}
		if t.Rightp != nil {
			printf(w, fmt.Sprintf("//t.Rightp: red"))
			printf(w, fmt.Sprintf("fill(255, 0, 0);"))
			printf(w, fmt.Sprintf("ellipse(%f, %f, 5.0, 5.0);",
				t.Rightp.X, t.Rightp.Y))
			printf(w, fmt.Sprintf("noFill();"))

			// the vertical through this point
			printf(w, fmt.Sprintf("stroke(255, 0, 0);"))
			printf(w, fmt.Sprintf("line(%f, %f, %f, %f);",
				t.Rightp.X, t.Rightp.Y, t.Rightp.X, 0.0))
			printf(w, fmt.Sprintf("noStroke();"))
		}
	}
	printf(w, fmt.Sprintf("}"))
	f.Close()
}
