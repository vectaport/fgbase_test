package main

import(
	"github.com/vectaport/flowgraph"
        "github.com/vectaport/flowgraph/regexp"
)

var teststrings = []string{
	"MMMapplesapplesapplesTAG",
	"applesapplesapples",
        "applesX",
        "orangesX",
	"apples",
	"oranges",
	"T",
	"",
}

func tbi(x flowgraph.Edge) flowgraph.Node {

        i := 0

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
                        j := i%(len(teststrings)+1)
                        if j<len(teststrings) {
  			    x.Val = teststrings[j]
                        } else {
                            x.Val = nil
                        }
                        i++
		})
	return node

}

func tbo(a flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
         
}

func main() {

	
	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(4, 4)
	
	n[0] = tbi(e[0])
        n[1] = regexp.FuncRegexp(e[0], e[1], e[2], e[3])
	n[2] = regexp.FuncMatch(e[2], e[1], "apples")
        n[3] = tbo(e[3])
	
	flowgraph.RunAll(n)

}
