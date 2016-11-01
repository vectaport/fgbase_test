package main

import (
        "fmt"
	"crypto/rand"

	"golang.org/x/crypto/nacl/box"

	"github.com/vectaport/flowgraph"
)

func tbi(x flowgraph.Edge) flowgraph.Node {

	node := flowgraph.MakeNode("tbi", nil, []*flowgraph.Edge{&x}, nil, 
		func (n *flowgraph.Node) {
			x.Val = fmt.Sprintf("XXX %d", n.Aux)
			n.Aux = n.Aux.(int) + 1
		})

	node.Aux = 0
	return node

}

func tbo(a flowgraph.Edge) flowgraph.Node {
	node := flowgraph.MakeNode("tbo", []*flowgraph.Edge{&a}, nil, nil, nil)
	return node
}


func main() {

	flowgraph.ConfigByFlag(map[string]interface{} {"sec": 2})

	publicKey1, privateKey1, _ := box.GenerateKey(rand.Reader)
	publicKey2, privateKey2, _ := box.GenerateKey(rand.Reader)
        nonce := [24]byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23}
        nonce2 := [24]byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23}

	e,n := flowgraph.MakeGraph(3,5)

	n[0] = tbi(e[0])
        n[1] = flowgraph.FuncEncrypt(e[0], e[1], privateKey1, publicKey2, nonce)
        n[2] = flowgraph.FuncDecrypt(e[1], e[2], privateKey2, publicKey1, nonce2)
        n[3] = tbo(e[2])
        n[4] = tbo(e[1])
	
	flowgraph.RunAll(n)

}

