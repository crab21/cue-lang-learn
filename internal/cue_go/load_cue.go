package cuego

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func cueLoad() {
	const config = `
msg:   "Hello \(place)!"
place: string | *"world" // "world" is the default.
`

	var r cue.Runtime

	instance, _ := r.Compile("", config)
	str, _ := instance.Lookup("msg").String()

	fmt.Println(str)

	c := cuecontext.New()
	c.BuildExpr()

}
