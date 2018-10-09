package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/nacl/box"

	"github.com/vectaport/fgbase"
)

func tbi(x fgbase.Edge) fgbase.Node {

	node := fgbase.MakeNode("tbi", nil, []*fgbase.Edge{&x}, nil,
		func(n *fgbase.Node) error {
			x.DstPut(fmt.Sprintf("XXX %d", n.Aux))
			n.Aux = n.Aux.(int) + 1
			return nil
		})

	node.Aux = 0
	return node

}

func tbo(a fgbase.Edge) fgbase.Node {
	node := fgbase.MakeNode("tbo", []*fgbase.Edge{&a}, nil, nil, nil)
	return node
}

func main() {

	fgbase.ConfigByFlag(map[string]interface{}{"sec": 2})

	publicKey1, privateKey1, _ := box.GenerateKey(rand.Reader)
	publicKey2, privateKey2, _ := box.GenerateKey(rand.Reader)
	nonce := [24]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	nonce2 := [24]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

	e, n := fgbase.MakeGraph(3, 5)

	n[0] = tbi(e[0])
	n[1] = fgbase.FuncEncrypt(e[0], e[1], privateKey1, publicKey2, nonce)
	n[2] = fgbase.FuncDecrypt(e[1], e[2], privateKey2, publicKey1, nonce2)
	n[3] = tbo(e[2])
	n[4] = tbo(e[1])

	fgbase.RunAll(n)

}
