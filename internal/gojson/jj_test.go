package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCueLoad(t *testing.T) {
	user := Users{"goog"}
	g, e := json.Marshal(user)
	if e != nil {
		t.Errorf(e.Error())
	}
	fmt.Println(string(g))
	e = json.Unmarshal([]byte("{\"od\":\"googfdsjfsadk\"}"), &user)
	if e != nil {
		t.Errorf(e.Error())
	}
	fmt.Println(user.Name)
}
