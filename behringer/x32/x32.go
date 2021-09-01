// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package x32

import (
	"fmt"
	"io"

	"github.com/goaudiovideo/osc"
)

// Mixer models a Behringer X32 mixer that can be controlled using Open Sound
// Control (OSC).
type Mixer struct {
	w io.Writer
}

// NewMixer creates a new Mixer using the given writer.
func NewMixer(w io.Writer) Mixer {
	return Mixer{
		w: w,
	}
}

// Write implements the Writer interface for Mixer.
func (m Mixer) Write(p []byte) (int, error) {
	return m.w.Write(p)
}

// MuteChannel mutes the given channel.
func (m Mixer) MuteChannel(ch int) error {
	addr := fmt.Sprintf("/ch/%2d/mix/on", ch)
	msg, err := osc.Message(addr, "i", 0)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}

// UnmuteChannel unmutes the given channel.
func (m Mixer) UnmuteChannel(ch int) error {
	addr := fmt.Sprintf("/ch/%2d/mix/on", ch)
	msg, err := osc.Message(addr, "i", 1)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}
