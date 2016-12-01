package main

import(	
	"github.com/vectaport/flowgraph"
        "github.com/vectaport/flowgraph/regexp"
)

var teststrings = []string{
	"test string",
	"test",
	"wrong",
}

func tbi(x flowgraph.Edge) flowgraph.Node {
	
        i := 0
	
	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x},
		func (n *flowgraph.Node) bool {
			return i<=len(teststrings) && n.DefaultRdyFunc()
		},
		func (n *flowgraph.Node) { 
                        if i<len(teststrings) {
				x.Val = teststrings[i]
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
	
	e,n := flowgraph.MakeGraph(3,3)

	e[2].Const("test")
	
	n[0] = tbi(e[0])
	n[1] = regexp.FuncMatch(e[0], e[2], e[1])
        n[2] = tbo(e[1])
	
	flowgraph.RunAll(n)
	
}
