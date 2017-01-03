package main

import(
	"github.com/vectaport/flowgraph"
        "github.com/vectaport/flowgraph/regexp"
)

var teststrings = []string{
	"test string",
	"apples",
	"oranges",
	"",
	"T",
        "applesX",
        "orangesX",
}

func tbi(x flowgraph.Edge) flowgraph.Node {

        i := 0

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x},
		func (n *flowgraph.Node) bool {
			return i<=len(teststrings) && n.DefaultRdyFunc()
		},
		func (n *flowgraph.Node) { 
                        if i<len(teststrings) {
				x.DstPut(regexp.Search{Curr:teststrings[i],Orig:teststrings[i]})
                        } else {
				if i==len(teststrings) {
					x.DstPut(regexp.Search{})
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

	e,n := flowgraph.MakeGraph(6,5)

	e[4].Const("apples")
	e[5].Const("oranges")
	
	n[0] = tbi(e[0])
	n[1] = regexp.FuncMatch(e[0], e[4], e[1])
	n[2] = regexp.FuncMatch(e[0], e[5], e[2])
        n[3] = regexp.FuncBar(e[1], e[2], e[3], true)
        n[4] = tbo(e[3])

	flowgraph.RunAll(n)

}
