package main

import(
	"bufio"
	"os"
	
	"github.com/vectaport/flowgraph"
        "github.com/vectaport/flowgraph/regexp"
)

var teststrings = []string{
	"axyz",
	"dxyz",
}

/*
var variants = []string{
    "agggtaaa|tttaccct",
    "[cgt]gggtaaa|tttaccc[acg]",
    "a[act]ggtaaa|tttacc[agt]t",
    "ag[act]gtaaa|tttac[agt]ct",
    "agg[act]taaa|ttta[agt]cct",
    "aggg[acg]aaa|ttt[cgt]ccct",
    "agggt[cgt]aa|tt[acg]accct",
    "agggta[cgt]a|t[acg]taccct",
    "agggtaa[cgt]|[acg]ttaccct",
}
*/

func check(e error) {
	if e != nil {
		flowgraph.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}
		
func tbi(dnstreq flowgraph.Edge, newmatch flowgraph.Edge) flowgraph.Node {

        i := 0

	f, err := os.Open("fasta.txt")
	check(err)
	r := bufio.NewReader(f)
	
	prev := make(map[string]string)

	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&dnstreq}, []*flowgraph.Edge{&newmatch},
		func (n *flowgraph.Node) bool {
			return dnstreq.SrcRdy(n) || newmatch.DstRdy(n) && i <= flowgraph.ChannelSize
		},
		func (n *flowgraph.Node) {
			if dnstreq.SrcRdy(n) {
				match := dnstreq.SrcGet().(regexp.Search)
				match.Curr = prev[match.Orig][1:]
				prev[match.Orig] = match.Curr
				newmatch.DstPut(match)
				i--
				return
			}
			xv,err := r.ReadString('\n')
			if err!=nil {
				newmatch.DstPut(regexp.Search{})
			} else {
				newmatch.DstPut(regexp.Search{Orig:xv, Curr:xv, State:regexp.Live})
			}
                        i++
		})
	return node
	
}

func tbo(oldmatch flowgraph.Edge, dnstreq flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&oldmatch}, []*flowgraph.Edge{&dnstreq}, nil,
		func (n *flowgraph.Node) {
			oldmatch.Flow = true
			dnstreq.DstPut(regexp.Search{State:regexp.Done}) // echo back
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
	test0
	test1
	subsrc0
	subsrc1
	edgeNum
)

var edgeNames []string = []string {
	"newmatch",
	"subsrc",
	"dnstreq",
	"oldmatch",
	"subdst",
	"upstreq",
	"test0",
	"test1",
	"subsrc0",
	"subsrc1",
}

func main() {
	
	
	flowgraph.ConfigByFlag(nil)
	
	e,n := flowgraph.MakeGraph(int(edgeNum),6)
	flowgraph.NameEdges(e,edgeNames)

	e[test0].Const("AGGGTAAA")
	e[test1].Const("TTTACCCT")
	
	n[0] = tbi(e[upstreq], e[newmatch])
	n[1] = regexp.FuncRepeat(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq], 1, -1)
	n[2] = regexp.FuncMatch(e[subdst], e[test0], e[subsrc0])
	n[3] = regexp.FuncMatch(e[subdst], e[test1], e[subsrc1])
	n[4] = regexp.FuncBar(e[subsrc0], e[subsrc1], e[subsrc], false)
	n[5] = tbo(e[oldmatch], e[dnstreq])
	
	flowgraph.RunAll(n)
	
}
