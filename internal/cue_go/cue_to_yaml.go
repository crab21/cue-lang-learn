package cuego

import (
	"encoding/json"
	"fmt"

	"cuelang.org/go/cue/cuecontext"
)

func CueToYaml() {
	const config = `import "encoding/yaml"

	configMap: data: "point.yaml":
	yaml.Marshal({
	x: 4.5
	y: 2.34
	})`
	ctx := cuecontext.New()
	instance := ctx.CompileString(config)
	s, e := instance.Eval().MarshalJSON()
	if e != nil {
		panic(e)
	}

	var m map[string]interface{}
	e = json.Unmarshal([]byte(s), &m)
	if e != nil {
		panic(e)
	}

	fmt.Println(m)
}
