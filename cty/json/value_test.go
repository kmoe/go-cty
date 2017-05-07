package json

import (
	"fmt"
	"testing"

	"github.com/apparentlymart/go-cty/cty"
)

func TestValueJSONable(t *testing.T) {
	tests := []struct {
		Value cty.Value
		Type  cty.Type
		Want  string
	}{
		// Primitives
		{
			cty.StringVal("hello"),
			cty.String,
			`"hello"`,
		},
		{
			cty.StringVal(""),
			cty.String,
			`""`,
		},
		{
			cty.StringVal("15"),
			cty.Number,
			`15`,
		},
		{
			cty.StringVal("true"),
			cty.Bool,
			`true`,
		},
		{
			cty.StringVal("1"),
			cty.Bool,
			`true`,
		},
		{
			cty.NullVal(cty.String),
			cty.String,
			`null`,
		},
		{
			cty.NumberIntVal(2),
			cty.Number,
			`2`,
		},
		{
			cty.NumberFloatVal(2.5),
			cty.Number,
			`2.5`,
		},
		{
			cty.NumberIntVal(5),
			cty.String,
			`"5"`,
		},
		{
			cty.True,
			cty.Bool,
			`true`,
		},
		{
			cty.False,
			cty.Bool,
			`false`,
		},
		{
			cty.True,
			cty.String,
			`"true"`,
		},

		// Lists
		{
			cty.ListVal([]cty.Value{cty.True, cty.False}),
			cty.List(cty.Bool),
			`[true,false]`,
		},
		{
			cty.ListValEmpty(cty.Bool),
			cty.List(cty.Bool),
			`[]`,
		},
		{
			cty.ListVal([]cty.Value{cty.True, cty.False}),
			cty.List(cty.String),
			`["true","false"]`,
		},

		// Sets
		{
			cty.SetVal([]cty.Value{cty.True, cty.False}),
			cty.Set(cty.Bool),
			`[false,true]`,
		},
		{
			cty.SetValEmpty(cty.Bool),
			cty.Set(cty.Bool),
			`[]`,
		},

		// Encoding into dynamic produces type information wrapper
		{
			cty.True,
			cty.DynamicPseudoType,
			`{"value":true,"type":"bool"}`,
		},
		{
			cty.StringVal("hello"),
			cty.DynamicPseudoType,
			`{"value":"hello","type":"string"}`,
		},
		{
			cty.NumberIntVal(5),
			cty.DynamicPseudoType,
			`{"value":5,"type":"number"}`,
		},
		{
			cty.ListVal([]cty.Value{cty.True, cty.False}),
			cty.DynamicPseudoType,
			`{"value":[true,false],"type":["list","bool"]}`,
		},
		{
			cty.ListVal([]cty.Value{cty.True, cty.False}),
			cty.List(cty.DynamicPseudoType),
			`[{"value":true,"type":"bool"},{"value":false,"type":"bool"}]`,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%#v to %#v", test.Value, test.Type), func(t *testing.T) {
			gotBuf, err := Marshal(test.Value, test.Type)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got := string(gotBuf)

			if got != test.Want {
				t.Errorf(
					"wrong serialization\nvalue: %#v\ntype:  %#v\ngot:   %s\nwant:  %s",
					test.Value, test.Type, got, test.Want,
				)
			}
		})
	}
}