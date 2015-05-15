package main

import (
	"time"

	"github.com/vectaport/flowgraph"
)

var images = []string{"airplane.jpg", "fruits.jpg", "pic1.png", "pic3.png", "pic5.png", "stuff.jpg",
	"baboon.jpg", "lena.jpg", "pic2.png", "pic4.png", "pic6.png"}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			x.Val = "../../lazywei/go-opencv/images/"+images[n.Cnt%int64(len(images))]
		})
	return node
}

func main() {

	flowgraph.TraceLevel = flowgraph.V

	e,n := flowgraph.MakeGraph(1,2)
 
	n[0] = tbi(e[0])
	n[1] = flowgraph.FuncDisplay(e[0])

	flowgraph.RunAll(n, time.Second)

}

