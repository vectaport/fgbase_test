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
			// if i > len(teststrings) { x.NoOut = true; return }
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

	e,n := flowgraph.MakeGraph(6, 6)
	
	n[0] = tbi(e[0])
        n[1] = regexp.FuncStar(e[0], e[4], e[1], e[5])
	n[2] = regexp.FuncMatch(e[1], e[2], "apples")
	n[3] = regexp.FuncMatch(e[1], e[3], "oranges")
        n[4] = regexp.FuncBar(e[2], e[3], e[4], false)
        n[5] = tbo(e[5])

	flowgraph.RunAll(n)

}
