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

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x},
		func (n *flowgraph.Node) bool {
			return i<=len(teststrings) && n.DefaultRdyFunc()
		},
		func (n *flowgraph.Node) { 
                        if i<len(teststrings) {
				x.Val = regexp.Regexp{Orig:teststrings[i], Curr:teststrings[i]}
                        } else {
				if i==len(teststrings) {
					x.Val = nil
				} else {
					x.NoOut = false
				}
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

	e,n := flowgraph.MakeGraph(5, 4)

	e[4].Const("apples")
	
	n[0] = tbi(e[0])
        n[1] = regexp.FuncRegexp(e[0], e[1], e[2], e[3])
	n[2] = regexp.FuncMatch(e[2], e[4], e[1])
        n[3] = tbo(e[3])
	
	flowgraph.RunAll(n)

}
