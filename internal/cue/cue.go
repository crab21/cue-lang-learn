package cue

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

const val = `
i: int
s: "hello"
`

func CueFunc() {
	CueInit()
	// loadCue()
}

func CueInit() {
	var (
		c *cue.Context
		v cue.Value
	)

	// create a context
	c = cuecontext.New()

	// compile some CUE into a Value
	v = c.CompileString(val)

	// print the value
	fmt.Println(v.String())
}

func loadCue() {
	ctx := cuecontext.New()

	// The entrypoints are the same as the files you'd specify at the command line
	entrypoints := []string{"/Users/gogoowang/workspace/cue-lang-learn/internal/cue/hello.cue"}

	// Load Cue files into Cue build.Instances slice
	// the second arg is a configuration object, we'll see this later
	bis := load.Instances(entrypoints, nil)

	// Loop over the instances, checking for errors and printing
	for _, bi := range bis {
		// check for errors on the instance
		// these are typically parsing errors
		if bi.Err != nil {
			fmt.Println("Error during load:", bi.Err)
			continue
		}

		// Use cue.Context to turn build.Instance to cue.Instance
		value := ctx.BuildInstance(bi)
		if value.Err() != nil {
			fmt.Println("Error during build:", value.Err())
			continue
		}

		// print the error
		fmt.Println("root value:", value)

		// Validate the value
		err := value.Validate()
		if err != nil {
			fmt.Println("Error during validate:", err)
			continue
		}
	}
}
