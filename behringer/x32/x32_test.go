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

func TestNameChannel(t *testing.T) {
	var tests = []struct {
		channel     int
		name        string
		want        string
		expectError bool
	}{
		{0, "foo", "/ch/00/config/name\x00\x00,s\x00\x00foo\x00", true},
		{1, "foo", "/ch/01/config/name\x00\x00,s\x00\x00foo\x00", false},
		{1, "badTooLongName", "/ch/01/config/name\x00\x00,s\x00\x00foo\x00", true},
		{32, "foo", "/ch/32/config/name\x00\x00,s\x00\x00foo\x00", false},
		{33, "foo", "/ch/33/config/name\x00\x00,s\x00\x00foo\x00", true},
	}
	for _, test := range tests {
		name := fmt.Sprintf("ch%02d_%s", test.channel, test.name)
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			mixer := NewMixer(&b)
			err := mixer.NameChannel(test.channel, test.name)
			if test.expectError {
				if err == nil {
					t.Errorf("expected error naming channel %d", test.channel)
				}
			} else {
				if err != nil {
					t.Errorf("error naming channel %d: %s", test.channel, err)
				}
				got := b.String()
				if got != test.want {
					t.Errorf("\t got = %x\n\t\t\twant = %x", got, test.want)
				}
			}
		})
	}
}

func TestSetChannelIcon(t *testing.T) {
	var tests = []struct {
		channel     int
		icon        Icon
		want        string
		expectError bool
	}{
		{0, BassDrum, "/ch/00/config/icon\x00\x00,i\x00\x00\x00\x00\x00\x03", true},
		{1, BassDrum, "/ch/01/config/icon\x00\x00,i\x00\x00\x00\x00\x00\x03", false},
		{1, BassDrum, "/ch/01/config/icon\x00\x00,i\x00\x00\x00\x00\x00\x03", false},
		{31, Laptop, "/ch/31/config/icon\x00\x00,i\x00\x00\x00\x00\x00\x3e", false},
		{32, Cymbal, "/ch/32/config/icon\x00\x00,i\x00\x00\x00\x00\x00\x0a", false},
		{33, Cymbal, "/ch/32/config/icon\x00\x00,i\x00\x00\x00\x00\x00\x0a", true},
	}
	for _, test := range tests {
		name := fmt.Sprintf("ch%02d_%s", test.channel, test.icon)
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			mixer := NewMixer(&b)
			err := mixer.SetChannelIcon(test.channel, test.icon)
			if test.expectError {
				if err == nil {
					t.Errorf("expected error setting icon for channel %d", test.channel)
				}
			} else {
				if err != nil {
					t.Errorf("error setting icon for channel %d: %s", test.channel, err)
				}
				got := b.String()
				if got != test.want {
					t.Errorf("\t got = %x\n\t\t\twant = %x", got, test.want)
				}
			}
		})
	}
}

func TestSetChannelColor(t *testing.T) {
	var tests = []struct {
		channel     int
		color       Color
		want        string
		expectError bool
	}{
		{0, RedText, "/ch/00/config/color\x00,i\x00\x00\x00\x00\x00\x01", true},
		{1, RedText, "/ch/01/config/color\x00,i\x00\x00\x00\x00\x00\x01", false},
		{32, BlueBackground, "/ch/32/config/color\x00,i\x00\x00\x00\x00\x00\x0c", false},
		{33, RedText, "/ch/33/config/color\x00,i\x00\x00\x00\x00\x00\x01", true},
	}
	for _, test := range tests {
		name := fmt.Sprintf("ch%02d_%s", test.channel, test.color)
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			mixer := NewMixer(&b)
			err := mixer.SetChannelColor(test.channel, test.color)
			if test.expectError {
				if err == nil {
					t.Errorf("expected error setting color for channel %d", test.channel)
				}
			} else {
				if err != nil {
					t.Errorf("error setting color for channel %d: %s", test.channel, err)
				}
				got := b.String()
				if got != test.want {
					t.Errorf("\t got = %x\n\t\t\twant = %x", got, test.want)
				}
			}
		})
	}
}

func TestConvertDBToDecimal(t *testing.T) {
	var tests = []struct {
		given float64
		want  float64
	}{
		{-150.0, 0.0},
		{-90.0, 0.0},
		{-78.0, 0.025},
		{-60.0, 0.0625},
		{-42.0, 0.175},
		{-30.0, 0.25},
		{-20.0, 0.375},
		{-10.0, 0.50},
		{0.0, 0.75},
		{10.0, 1.0},
		{20.0, 1.0},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%.3fdB", test.given)
		t.Run(name, func(t *testing.T) {
			if got := dbLevelToDecimal(test.given); got != test.want {
				t.Errorf("\t got = %.4f\n\t\t\twant = %.4f", got, test.want)
			}
		})
	}
}

func TestConvertDecimalToDB(t *testing.T) {
	var tests = []struct {
		given float64
		want  float64
	}{
		{-10.00, -90.0},
		{0.0000, -90.0},
		{0.0250, -78.0},
		{0.0625, -60.0},
		{0.1750, -42.0},
		{0.2500, -30.0},
		{0.3750, -20.0},
		{0.5000, -10.0},
		{0.7500, +0.00},
		{1.0000, +10.0},
		{1.1000, +10.0},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%.3fdB", test.given)
		t.Run(name, func(t *testing.T) {
			if got := decimalToDBLevel(test.given); got != test.want {
				t.Errorf("\t got = %.4f\n\t\t\twant = %.4f", got, test.want)
			}
		})
	}
}
