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

	f, err := os.Open("fasta.huge.txt")
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
			n.Tracef("NEW MATCH FROM UPSTREAM %d\n", i+1);
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
	subsrcQ
	subsrcBB
	subsrcCC
	subsrcDD
	subsrcEE
	subsrcFF
	subsrcGG
	subsrcHH
	subsrcII
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
	"subsrcQ",
	"subsrcAA",
	"subsrcBB",
	"subsrcCC",
	"subsrcDD",
	"subsrcEE",
	"subsrcFF",
	"subsrcGG",
	"subsrcHH",
	"subsrcII",
}

func main() {
	
	
	flowgraph.ConfigByFlag(nil)
	
	e,n := flowgraph.MakeGraph(int(edgeNum), 47)
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
	
	n[fi()] = flowgraph.FuncPrint(e[subsrcA], e[subsrcJ], "pattern 1: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcJ], e[subsrcB], e[subsrcBB] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcBB], e[subsrcK], "pattern 2: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcK], e[subsrcC], e[subsrcCC] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcCC], e[subsrcL], "pattern 3: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcL], e[subsrcD], e[subsrcDD] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcDD], e[subsrcM], "pattern 4: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcM], e[subsrcE], e[subsrcEE] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcEE], e[subsrcN], "pattern 5: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcN], e[subsrcF], e[subsrcFF] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcFF], e[subsrcO], "pattern 6: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcO], e[subsrcG], e[subsrcGG] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcGG], e[subsrcP], "pattern 7: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcP], e[subsrcH], e[subsrcHH] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcHH], e[subsrcQ], "pattern 8: %v\n")
	
	n[fi()] = flowgraph.FuncRdy(e[subsrcQ], e[subsrcI], e[subsrcII] )
	n[fi()] = flowgraph.FuncPrint(e[subsrcII], e[subsrc], "pattern 9: %v\n")
	
	n[fi()] = tbo(e[oldmatch], e[dnstreq])
	
	flowgraph.RunAll(n)

}
