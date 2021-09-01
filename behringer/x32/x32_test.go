// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package x32

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMuteChannel(t *testing.T) {
	var tests = []struct {
		channel     int
		want        string
		expectError bool
	}{
		{0, "/ch/01/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x00", true},
		{1, "/ch/01/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x00", false},
		{2, "/ch/02/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x00", false},
		{32, "/ch/32/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x00", false},
		{33, "/ch/01/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x00", true},
	}
	for _, test := range tests {
		name := fmt.Sprintf("ch%2d", test.channel)
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			mixer := NewMixer(&b)
			err := mixer.MuteChannel(test.channel)
			if test.expectError {
				if err == nil {
					t.Errorf("expected error muting channel %d", test.channel)
				}
			} else {
				if err != nil {
					t.Errorf("error muting channel %d: %s", test.channel, err)
				}
				got := b.String()
				if got != test.want {
					t.Errorf("\t got = %x\n\t\t\twant = %x", got, test.want)
				}
			}
		})
	}

}

func TestUnmuteChannel(t *testing.T) {
	var tests = []struct {
		channel     int
		want        string
		expectError bool
	}{
		{0, "/ch/01/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x01", true},
		{1, "/ch/01/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x01", false},
		{2, "/ch/02/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x01", false},
		{32, "/ch/32/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x01", false},
		{33, "/ch/01/mix/on\x00\x00\x00,i\x00\x00\x00\x00\x00\x01", true},
	}
	for _, test := range tests {
		name := fmt.Sprintf("ch%2d", test.channel)
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			mixer := NewMixer(&b)
			err := mixer.UnmuteChannel(test.channel)
			if test.expectError {
				if err == nil {
					t.Errorf("expected error unmuting channel %d", test.channel)
				}
			} else {
				if err != nil {
					t.Errorf("error unmuting channel %d: %s", test.channel, err)
				}
				got := b.String()
				if got != test.want {
					t.Errorf("\t got = %x\n\t\t\twant = %x", got, test.want)
				}
			}
		})
	}

}

func TestMuteMain(t *testing.T) {
	want := "/main/st/mix/on\x00,i\x00\x00\x00\x00\x00\x00"
	var b bytes.Buffer
	mixer := NewMixer(&b)
	if err := mixer.MuteMain(); err != nil {
		t.Errorf("error muting main: %s", err)
	}
	got := b.String()
	if got != want {
		t.Errorf("\t got = %x\n\t\twant = %x", got, want)
	}
}

func TestUnmuteMain(t *testing.T) {
	want := "/main/st/mix/on\x00,i\x00\x00\x00\x00\x00\x01"
	var b bytes.Buffer
	mixer := NewMixer(&b)
	if err := mixer.UnmuteMain(); err != nil {
		t.Errorf("error unmuting main: %s", err)
	}
	got := b.String()
	if got != want {
		t.Errorf("\t got = %x\n\t\twant = %x", got, want)
	}
}
