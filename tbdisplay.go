package main

import (
	"flag"

	"github.com/lazywei/go-opencv/opencv"
	"github.com/vectaport/fgbase"
	"github.com/vectaport/fgbase/imglab"
)

var images = []string{"airplane.jpg", "fruits.jpg", "pic1.png", "pic3.png", "pic5.png", "stuff.jpg",
	"baboon.jpg", "lena.jpg", "pic2.png", "pic4.png", "pic6.png"}

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error {
			filename := "../../lazywei/go-opencv/images/" + images[n.Cnt%int64(len(images))]
			n.Tracef("Loading %s\n", filename)
			x.DstPut(opencv.LoadImage(filename))
			return nil
		})
	return node
}

func main() {

	testp := flag.Bool("test", false, "test mode")
	fgbase.ConfigByFlag(nil)
	test := *testp

	var quitChan chan struct{}
	if !test {
		quitChan = make(chan struct{})
		fgbase.RunTime = 0
	}

	e, n := fgbase.MakeGraph(1, 2)

	n[0] = tbi(e[0])
	n[1] = imglab.FuncDisplay(e[0], quitChan)

	fgbase.RunAll(n)

	if !test {
		<-quitChan
	}
}
