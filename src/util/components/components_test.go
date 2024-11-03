package components_test

import (
	"testing"

	. "github.com/ValiantChip/gopck/src/util/components"
)

type TestComponent struct {
	key   string
	value any
}

func (t TestComponent) Key() string {
	return t.key
}

func (t TestComponent) Value() any {
	return t.value
}

var componentTests = []struct {
	in  []Component
	out string
}{
	{
		in:  []Component{TestComponent{key: "foo", value: true}, TestComponent{key: "bar", value: int8(12)}},
		out: "[foo=true,bar=12]",
	},
	{
		in:  []Component{TestComponent{key: "foo", value: "baz"}, TestComponent{key: "bar", value: float64(12.34)}},
		out: "[foo=\"baz\",bar=12.34]",
	},
	{
		in:  []Component{TestComponent{key: "foo", value: []any{1, 2, 3}}, TestComponent{key: "bar", value: map[string]any{"a": 1, "b": false}}},
		out: "[foo=[1,2,3],bar={\"a\":1,\"b\":false}]",
	},
	{
		in:  []Component{TestComponent{key: "foo", value: int64(132424)}},
		out: "[foo=132424]",
	},
	{
		in:  []Component{TestComponent{key: "foo", value: int32(132424)}, TestComponent{key: "bar", value: 56}},
		out: "[foo=132424,bar=56]",
	},
	{
		in:  []Component{},
		out: "[]",
	},
}

func TestParse(t *testing.T) {
	for _, tc := range componentTests {
		t.Run(tc.out, func(t *testing.T) {
			s, err := Parse(tc.in)
			if err != nil {
				t.Error(err)
			}

			if s != tc.out {
				t.Errorf("got %s, want %s", s, tc.out)
			}
		})
	}
}
