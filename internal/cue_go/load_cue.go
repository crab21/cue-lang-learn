package cuego

import (
	"fmt"

	"cuelang.org/go/cue"
)

func cueLoad() {
	const config = `
msg:   "Hello \(place)!"
place: string | *"world" // "world" is the default.
`

	var r cue.Runtime

	instance, _ := r.Compile("", config)
	str, _ := instance.Lookup("place").String()

	fmt.Println(str)

}
