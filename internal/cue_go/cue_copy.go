package cuego

import (
	"fmt"

	"cuelang.org/go/cue"
)

func CopyValueToGo() {
	type ab struct{ A, B int }

	type ccb struct {
		B []byte
	}
	type cc struct {
		B string
	}
	var r cue.Runtime

	var x ab
	var yb ccb
	var y cc

	i, _ := r.Compile("test", `{A: 2, B: 4}`)
	_ = i.Value().Decode(&x)
	fmt.Println(x)

	// correspond type: `byte`
	i, _ = r.Compile("test", `{B: 'foo'}`)
	_ = i.Value().Decode(&yb)
	fmt.Println(string(yb.B))

	// correspond type: `string`
	i, _ = r.Compile("test", `{B: "foo"}`)
	_ = i.Value().Decode(&y)
	fmt.Println(y)
}
