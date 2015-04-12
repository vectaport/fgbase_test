package main

import (
	"github.com/vectaport/flowgraph"
	"fmt"
	"math"
	"time"
	"reflect"
)

func promoteTest(a, b, x flowgraph.Edge) {
	
	for {
		aTmp := <- a.Data
		bTmp := <- b.Data
		fmt.Printf("%v,%v --> ", reflect.TypeOf(aTmp), reflect.TypeOf(bTmp))
		
		aBig,bBig,same := flowgraph.Promote(nil, aTmp, bTmp)
		
		fmt.Printf("%v,%v,%v\n", reflect.TypeOf(aBig), reflect.TypeOf(bBig), same);
		
		x.Data <- aBig
	}
	
	
}
func main() {

	a := flowgraph.MakeEdge("a",nil)
	b := flowgraph.MakeEdge("b",nil)
	x := flowgraph.MakeEdge("x",nil)

	go promoteTest(a, b, x)

  	var answer interface {}
	a.Data <- 512
	b.Data <- int8(4)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))
	
	a.Data <- byte(4)
	b.Data <- 512
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))
	
	a.Data <- byte(4)
	b.Data <- byte(100)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))
	
	a.Data <- "abcdef"
	b.Data <- byte(4)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- int8(8)
	b.Data <- uint32(4)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- 1 + 0i
	b.Data <- uint32(4)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- complex(float32(1),0)
	b.Data <- int64(4)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- float32(0)
	b.Data <- byte(0)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- uint64(math.MaxUint64)
	b.Data <- int64(-1)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- uint64(math.MaxUint64>>2)
	b.Data <- int64(-1)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	a.Data <- rune(33)
	b.Data <- int8(-1)
        answer = <- x.Data
	fmt.Printf("answer is %v of type %v\n\n", answer, reflect.TypeOf(answer))

	time.Sleep(time.Second)

}

