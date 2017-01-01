package main

import (
	"flag"

	"github.com/lazywei/go-opencv/opencv"
	"github.com/vectaport/flowgraph"
	"github.com/vectaport/flowgraph/imglab"
)

var images = []string{"airplane.jpg", "fruits.jpg", "pic1.png", "pic3.png", "pic5.png", "stuff.jpg",
	"baboon.jpg", "lena.jpg", "pic2.png", "pic4.png", "pic6.png"}

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
			filename := "../../lazywei/go-opencv/images/"+images[n.Cnt%int64(len(images))]
			n.Tracef("Loading %s\n", filename)
			x.DstPut(opencv.LoadImage(filename))
		})
	return node
}

func main() {

	
	testp := flag.Bool("test", false, "test mode")
	flowgraph.ConfigByFlag(nil)
	test := *testp

	var quitChan chan struct{}
	if !test {
		quitChan =make(chan struct{})
		flowgraph.RunTime = 0
	}

	e,n := flowgraph.MakeGraph(1,2)
 
	n[0] = tbi(e[0])
	n[1] = imglab.FuncDisplay(e[0], quitChan)

	flowgraph.RunAll(n)

	if !test {
		<- quitChan
	}
}

