// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package osc

import "testing"

func TestGoodMessages(t *testing.T) {
	var tests = []struct {
		addr    string
		typeTag string
		args    []interface{}
		want    []byte
	}{
		{"/info", "", nil, []byte("/info\x00\x00\x00,\x00\x00\x00")},
		{"/ch/01/config/name", "s", nil, []byte("/ch/01/config/name\x00\x00,s\x00\x00")},
		{"/ch/01/config/name", "s", nil, []byte("/ch/01/config/name\x00\x00,s\x00\x00")},
	}
	for _, test := range tests {
		got, err := Message(test.addr, test.typeTag, test.args)
		if err != nil {
			t.Errorf("Got unexpected error: %s", err)
		}
		if string(got) != string(test.want) {
			t.Errorf("got = %q / want = %q", got, test.want)
		}
	}
}

func TestNumPadBytes(t *testing.T) {
	var tests = []struct {
		given int
		want  int
	}{
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
	}
	for _, test := range tests {
		if got := numPadBytes(test.given); got != test.want {
			t.Errorf("given = %d / got = %d / want %d", test.given, got, test.want)
		}
	}
}
