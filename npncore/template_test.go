package npncore

import (
	"fmt"
	"testing"
)

var tests = []struct {
	Src  string
	Data Data
	Tgt  string
}{
	{Src: "a{b}c", Data: nil, Tgt: "abc"},                                  // Missing
	{Src: "a{b}c", Data: Data{"b": "x"}, Tgt: "axc"},                       // Basic
	{Src: "a{b}c", Data: Data{"b": "{foo}zz", "foo": "xx"}, Tgt: "axxzzc"}, // Recursive
	{Src: "a{b|default}c", Data: nil, Tgt: "adefaultc"},                    // Default
	{Src: "a{b|default}c", Data: Data{"b": "x"}, Tgt: "axc"},               // Skip default
}

// Tests the templating system
func TestMerge(test *testing.T) {
	for _, t := range tests {
		r, err := Merge(t.Src, t.Data)
		if err != nil {
			test.Error(err, "merge error for ["+t.Src+"]")
		}
		if r != t.Tgt {
			test.Error(fmt.Sprintf("merge expected [%v] but observed [%v]", t.Tgt, r))
		}
	}
}
