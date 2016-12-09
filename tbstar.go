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

func tbi(dnstreq flowgraph.Edge, newmatch flowgraph.Edge) flowgraph.Node {

        i := 0

	Prev:=make(map[string]string)

	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&dnstreq}, []*flowgraph.Edge{&newmatch},
		func (n *flowgraph.Node) bool {
			return dnstreq.SrcRdy(n) || newmatch.DstRdy(n) && i<len(teststrings)
		},
		func (n *flowgraph.Node) {
			if dnstreq.SrcRdy(n) {
				match := dnstreq.Val.(regexp.Search)
				match.Curr = Prev[match.Orig][1:]
				Prev[match.Orig] = match.Curr
				newmatch.Val = match
				return
			}
			dnstreq.NoOut = true
                        if i<len(teststrings) {
				newmatch.Val = regexp.Search{Orig:teststrings[i], Curr:teststrings[i], State:regexp.Live}
                        } else {
				if i==len(teststrings) {
					newmatch.Val = regexp.Search{}
				} else {
					newmatch.NoOut = true
				}
                        }
                        i++
			dnstreq.NoOut = true
		})
	return node

}

func tbo(oldmatch flowgraph.Edge, dnstreq flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&oldmatch}, []*flowgraph.Edge{&dnstreq}, nil,
		func (n *flowgraph.Node) {
			dnstreq.Val = regexp.Search{} // echo back
		})
	return node
         
}


type edgeCnt int
const (
	newmatch edgeCnt = iota
	subsrc
	dnstreq
	oldmatch
	subdst
	upstreq
	apples
	oranges
	applesmatch
	orangesmatch
	edgenum
)

var edgeNames []string = []string {
	"newmatch",
	"subsrc",
	"dnstreq",
	"oldmatch",
	"subdst",
	"upstreq",
	"apples",
	"oranges",
	"applesmatch",
	"orangesmatch",
}

	
func main() {

	flowgraph.ConfigByFlag(nil)

	e,n := flowgraph.MakeGraph(int(edgenum), 6)
	flowgraph.NameEdges(e,edgeNames)


	e[apples].Const("apples")
	e[oranges].Const("oranges")

	// (apples|oranges)*
	n[0] = tbi(e[upstreq], e[newmatch])
        n[1] = regexp.FuncStar(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq])
	n[2] = regexp.FuncMatch(e[subdst], e[apples], e[applesmatch])
	n[3] = regexp.FuncMatch(e[subdst], e[oranges], e[orangesmatch])
        n[4] = regexp.FuncBar(e[applesmatch], e[orangesmatch], e[subsrc], false)
        n[5] = tbo(e[oldmatch], e[dnstreq])

	flowgraph.RunAll(n)

}
