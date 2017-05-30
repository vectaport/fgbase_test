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
				match := dnstreq.SrcGet().(regexp.Search)
				if len(Prev[match.Orig])>2 {
  				         match.Curr = Prev[match.Orig][1:]
				         Prev[match.Orig] = match.Curr
				         newmatch.DstPut(match)
   				         return
			        }
				delete(Prev, match.Orig)
				return
			}
                        if i<len(teststrings) {
				newmatch.DstPut(regexp.Search{Orig:teststrings[i], Curr:teststrings[i], State:regexp.Live, ID:regexp.NextID()})
                        } else {
				if i==len(teststrings) {
					newmatch.DstPut(regexp.Search{ID:regexp.NextID()})
				}
                        }
                        i++
		})
	return node

}

func tbo(oldmatch flowgraph.Edge, dnstreq flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&oldmatch}, []*flowgraph.Edge{&dnstreq}, nil,
		func (n *flowgraph.Node) {
			match := oldmatch.SrcGet().(regexp.Search)
			match.State = regexp.Done
			dnstreq.DstPut(match) // echo back
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
	edgeNum
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

	e,n := flowgraph.MakeGraph(int(edgeNum), 6)
	flowgraph.NameEdges(e,edgeNames)


	e[apples].Const("apples")
	e[oranges].Const("oranges")

	// (apples|oranges)*banana
	n[0] = tbi(e[upstreq], e[newmatch])
        n[1] = regexp.FuncRepeat(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq], 0, -1)
	n[2] = regexp.FuncMatch(e[subdst], e[apples], e[applesmatch], false)
	n[3] = regexp.FuncMatch(e[subdst], e[oranges], e[orangesmatch], false)
        n[4] = regexp.FuncBar(e[applesmatch], e[orangesmatch], e[subsrc], false)
        n[5] = tbo(e[oldmatch], e[dnstreq])

	flowgraph.RunAll(n)

}
