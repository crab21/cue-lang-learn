package cuego

import (
	"fmt"

	"cuelang.org/go/cue"
)

func ModifyCUEValues() {

	const config = `
	msg:   "Hello \(place)!"
	place: string | *"world" // "world" is the default.
	`

	var r cue.Runtime

	instance, _ := r.Compile("", config)
	str, _ := instance.Lookup("place").String()

	inst, _ := instance.Fill("you", "place")
	str, _ = inst.Lookup("msg").String()

	fmt.Println(str)
}
