// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package osc

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGoodMessages(t *testing.T) {
	var args1 []interface{}
	args1 = append(args1, "name")
	var args2 []interface{}
	args2 = append(args2, float32(0.4648))
	var args3 []interface{}
	args3 = append(args3, 1)
	var tests = []struct {
		name    string
		addr    string
		typeTag string
		args    []interface{}
		want    []byte
	}{
		{"info", "/info", "", nil, []byte("/info\x00\x00\x00,\x00\x00\x00")},
		{
			"config ch1 name no args", "/ch/01/config/name", "s", nil,
			[]byte("/ch/01/config/name\x00\x00,s\x00\x00"),
		},
		{
			"config ch1 name", "/ch/01/config/name", "s", args1,
			[]byte("/ch/01/config/name\x00\x00,s\x00\x00name\x00\x00\x00\x00"),
		},
		{
			"ch1 freq", "/ch/01/eq/1/q", "f", args2,
			[]byte("/ch/01/eq/1/q\x00\x00\x00,f\x00\x00\x3e\xed\xfa\x44"),
		},
		{
			"ch1 gate mode", "/ch/01/gate/mode", "i", args3,
			[]byte("/ch/01/gate/mode\x00\x00\x00\x00,i\x00\x00\x00\x00\x00\x01"),
		},
		{
			"ch2 gate mode", "/ch/02/gate/mode", "i", args3,
			[]byte("/ch/02/gate/mode\x00\x00\x00\x00,i\x00\x00\x00\x00\x00\x01"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Message(test.addr, test.typeTag, test.args...)
			if err != nil {
				t.Errorf("Got unexpected error: %s", err)
			}
			if string(got) != string(test.want) {
				t.Errorf("\t got = %q\n\t\t\twant = %q", got, test.want)
			}
		})
	}
}

func TestNumZeroBytes(t *testing.T) {
	var tests = []struct {
		given int
		want  int
	}{
		{0, 0},
		{1, 3},
		{2, 2},
		{3, 1},
		{4, 0},
		{5, 3},
		{6, 2},
		{7, 1},
		{8, 0},
		{9, 3},
		{10, 2},
		{11, 1},
		{12, 0},
		{13, 3},
		{14, 2},
		{15, 1},
		{16, 0},
		{17, 3},
	}
	for _, test := range tests {
		name := fmt.Sprintf("len %d", test.given)
		t.Run(name, func(t *testing.T) {
			if got := numZeroBytes(test.given); got != test.want {
				t.Errorf("\t got = %d\n\t\t\twant = %d", got, test.want)
			}
		})
	}
}

func TestEncodeFloat32(t *testing.T) {
	var tests = []struct {
		given float32
		want  string
	}{
		{0.0000, "00000000"},
		{0.0010, "3a83126f"},
		{0.0020, "3b03126f"},
		{0.2650, "3e87ae14"},
		{0.4648, "3eedfa44"},
		{0.5000, "3f000000"},
		{0.7500, "3f400000"},
		{0.8250, "3f533333"},
	}
	for _, test := range tests {
		name := fmt.Sprintf("float_%f", test.given)
		t.Run(name, func(t *testing.T) {
			h, err := hex.DecodeString(test.want)
			if err != nil {
				t.Errorf("unexepcted error decoding hex string %s: %s", test.want, err)
			}
			got, err := encodeFloat32(test.given)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if string(got) != string(h) {
				t.Errorf("\t got = %x\n\t\t\twant = %x", got, h)
			}
		})
	}
}

func TestEncodeInt(t *testing.T) {
	var tests = []struct {
		given int
		want  string
	}{
		{0, "00000000"},
		{1, "00000001"},
		{2, "00000002"},
		{3, "00000003"},
		{4, "00000004"},
	}
	for _, test := range tests {
		name := fmt.Sprintf("integer_%d", test.given)
		t.Run(name, func(t *testing.T) {
			h, err := hex.DecodeString(test.want)
			if err != nil {
				t.Errorf("unexpected error decoding hex string %s: %s", test.want, err)
			}
			got, err := encodeInt(test.given)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if string(got) != string(h) {
				t.Errorf("\t got = %x\n\t\t\twant = %x", got, h)
			}
		})
	}
}
