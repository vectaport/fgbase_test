package main

import(
	"github.com/vectaport/flowgraph"
        "github.com/vectaport/flowgraph/regexp"
)

var teststrings = []string{"test string", "test", "wrong"}

func tbi(x flowgraph.Edge) flowgraph.Node {

        i := 0

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) { 
                        j := i%(len(teststrings)+1)
                        if j<len(teststrings)-1 {
  			    x.Val = []byte(teststrings[j])
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

	e,n := flowgraph.MakeGraph(2,3)
 
	n[0] = tbi(e[0])
	n[1] = regexp.FuncMatch(e[0], e[1], "test")
        n[2] = tbo(e[1])

	flowgraph.RunAll(n)

}
