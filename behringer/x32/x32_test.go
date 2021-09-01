// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package x32

import (
	"fmt"
	"testing"
)

func TestMuteChannel(t *testing.T) {
	var tests = []struct {
		channel int
	}{
		{1},
		{2},
	}
	for _, test := range tests {
		name := fmt.Sprintf("ch%2d", test.channel)
		t.Run(name, func(t *testing.T) {
		})
	}

}
