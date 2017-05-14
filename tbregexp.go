package main

import(
	"bufio"
	"io"
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
    "AGGGTAAA|TTTACCCT",
    "[CGT]GGGTAAA|TTTACCC[ACG]",
    "A[ACT]GGTAAA|TTTACC[AGT]T",
    "AG[ACT]GTAAA|TTTAC[AGT]CT",
    "AGG[ACT]TAAA|TTTA[AGT]CCT",
!    "AGGG[ACG]AAA|TTT[CGT]CCCT",
    "AGGGT[CGT]AA|TT[ACG]ACCCT",
    "AGGGTA[CGT]A|T[ACG]TACCCT",
    "AGGGTAA[CGT]|[ACG]TTACCCT",
}

*/

func check(e error) {
	if e != nil {
		flowgraph.StderrLog.Printf("%v\n", e)
		os.Exit(1)
	}
}

var prev map[int64]string;
		
func tbi(dnstreq flowgraph.Edge, newmatch flowgraph.Edge) flowgraph.Node {

        i := 0
	done := false

	f, err := os.Open("fasta.txt")
	check(err)
	r := bufio.NewReader(f)
	
	prev = make(map[int64]string)

	node := flowgraph.MakeNode("tbi", []*flowgraph.Edge{&dnstreq}, []*flowgraph.Edge{&newmatch},
		func (n *flowgraph.Node) bool {
			return !done && (dnstreq.SrcRdy(n) || newmatch.DstRdy(n) && i <= flowgraph.ChannelSize)
		},
		func (n *flowgraph.Node) {
			if dnstreq.SrcRdy(n) {
				match := dnstreq.SrcGet().(regexp.Search)
				if match.State == regexp.Done {
				        n.Tracef("DONE REQUEST FROM DOWNSTREAM %d\n", i-1);
				        delete(prev, match.ID)
					i--
				        return
				}
				n.Tracef("LIVE REQUEST FROM DOWNSTREAM %d\n", i);
				match.Curr = prev[match.ID][1:]
				prev[match.ID] = match.Curr
				newmatch.DstPut(match)
				return
			}
			xv,err := r.ReadString('\n')
			if err == io.EOF {
			        n.Tracef("EOF\n")
				done = true
			        return
			}
			id := regexp.NextID()
		        prev[id] = xv
			newmatch.DstPut(regexp.Search{Orig:xv, Curr:xv, State:regexp.Live, ID:id})
                        i++
		})
	return node
	
}

func tbo(oldmatch flowgraph.Edge, dnstreq flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&oldmatch}, []*flowgraph.Edge{&dnstreq}, nil,
		func (n *flowgraph.Node) {
			o := oldmatch.SrcGet().(regexp.Search)
			o.State = regexp.Done
			dnstreq.DstPut(o) // echo back
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
	test2
	test3
	test4
	test5
	test6
	test7
	test8
	test9
	test10
	test11
	test12
	test13
	test14
	test15
	test16
	test17
	subsrc0
	subsrc1
	subsrc2
	subsrc3
	subsrc4
	subsrc5
	subsrc6
	subsrc7
	subsrc8
	subsrc9
	subsrc10
	subsrc11
	subsrc12
	subsrc13
	subsrc14
	subsrc15
	subsrc16
	subsrc17
	subsrcA
	subsrcB
	subsrcC
	subsrcD
	subsrcE
	subsrcF
	subsrcG
	subsrcH
	subsrcI
	subsrcJ
	subsrcK
	subsrcL
	subsrcM
	subsrcN
	subsrcO
	subsrcP
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
	"test2",
	"test3",
	"test4",
	"test5",
	"test6",
	"test7",
	"test8",
	"test9",
	"test10",
	"test11",
	"test12",
	"test13",
	"test14",
	"test15",
	"test16",
	"test17",
	"subsrc0",
	"subsrc1",
	"subsrc2",
	"subsrc3",
	"subsrc4",
	"subsrc5",
	"subsrc6",
	"subsrc7",
	"subsrc8",
	"subsrc9",
	"subsrc10",
	"subsrc11",
	"subsrc12",
	"subsrc13",
	"subsrc14",
	"subsrc15",
	"subsrc16",
	"subsrc17",
	"subsrcA",
	"subsrcB",
	"subsrcC",
	"subsrcD",
	"subsrcE",
	"subsrcF",
	"subsrcG",
	"subsrcH",
	"subsrcI",
	"subsrcJ",
	"subsrcK",
	"subsrcL",
	"subsrcM",
	"subsrcN",
	"subsrcO",
	"subsrcP",
}

func main() {
	
	
	flowgraph.ConfigByFlag(nil)
	
	e,n := flowgraph.MakeGraph(int(edgeNum), 38)
	flowgraph.NameEdges(e,edgeNames)


	// 1 match
        e[test0].Const("AGGGTAAA")
        e[test1].Const("TTTACCCT")
	
	// 0 match
	e[test2].Const("[CGT]GGGTAAA")
        e[test3].Const("TTTACCC[ACG]")
	
	// 0 match
	e[test4].Const("A[ACT]GGTAAA")
        e[test5].Const("TTTACC[AGT]T")
	
	// 0 match
	e[test6].Const("AG[ACT]GTAAA")
        e[test7].Const("TTTAC[AGT]CT")
	
	// 1 match
	e[test8].Const("AGG[ACT]TAAA")
        e[test9].Const("TTTA[AGT]CCT")
	
	// 0 match
	e[test10].Const("AGGG[ACG]AAA")
        e[test11].Const("TTT[CGT]CCCT")
	
	// 0 match
	e[test12].Const("AGGGT[CGT]AA")
        e[test13].Const("TT[ACG]ACCCT")
	
	// 0 match
	e[test14].Const("AGGGTA[CGT]A")
        e[test15].Const("T[ACG]TACCCT")
	
	// 2 match
	e[test16].Const("AGGGTAA[CGT]")
	e[test17].Const("[ACG]TTACCCT")

        i := 0
	fi := func() int {j := i; i++; return j}
	
	n[fi()] = tbi(e[upstreq], e[newmatch])
	n[fi()] = regexp.FuncRepeat(e[newmatch], e[subsrc], e[dnstreq], e[oldmatch], e[subdst], e[upstreq], 1, -1)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test0], e[subsrc0], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test1], e[subsrc1], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test2], e[subsrc2], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test3], e[subsrc3], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test4], e[subsrc4], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test5], e[subsrc5], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test6], e[subsrc6], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test7], e[subsrc7], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test8], e[subsrc8], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test9], e[subsrc9], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test10], e[subsrc10], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test11], e[subsrc11], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test12], e[subsrc12], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test13], e[subsrc13], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test14], e[subsrc14], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test15], e[subsrc15], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test16], e[subsrc16], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncMatch(e[subdst], e[test17], e[subsrc17], /* ignoreCase = */ true)
	n[fi()] = regexp.FuncBar(e[subsrc0], e[subsrc1], e[subsrcA], false)
	n[fi()] = regexp.FuncBar(e[subsrc2], e[subsrc3], e[subsrcB], false)
	n[fi()] = regexp.FuncBar(e[subsrc4], e[subsrc5], e[subsrcC], false)
	n[fi()] = regexp.FuncBar(e[subsrc6], e[subsrc7], e[subsrcD], false)
	n[fi()] = regexp.FuncBar(e[subsrc8], e[subsrc9], e[subsrcE], false)
	n[fi()] = regexp.FuncBar(e[subsrc10], e[subsrc11], e[subsrcF], false)
	n[fi()] = regexp.FuncBar(e[subsrc12], e[subsrc13], e[subsrcG], false)
	n[fi()] = regexp.FuncBar(e[subsrc14], e[subsrc15], e[subsrcH], false)
	n[fi()] = regexp.FuncBar(e[subsrc16], e[subsrc17], e[subsrcI], false)
	n[fi()] = regexp.FuncBar(e[subsrcA], e[subsrcB], e[subsrcJ], false)
	n[fi()] = regexp.FuncBar(e[subsrcC], e[subsrcD], e[subsrcK], false)
	n[fi()] = regexp.FuncBar(e[subsrcE], e[subsrcF], e[subsrcL], false)
	n[fi()] = regexp.FuncBar(e[subsrcG], e[subsrcH], e[subsrcM], false)
	n[fi()] = regexp.FuncBar(e[subsrcJ], e[subsrcK], e[subsrcN], false)
	n[fi()] = regexp.FuncBar(e[subsrcL], e[subsrcM], e[subsrcO], false)
	n[fi()] = regexp.FuncBar(e[subsrcN], e[subsrcO], e[subsrcP], false)
	n[fi()] = regexp.FuncBar(e[subsrcI], e[subsrcP], e[subsrc], false)
	n[fi()] = tbo(e[oldmatch], e[dnstreq])
	
	flowgraph.RunAll(n)

}
