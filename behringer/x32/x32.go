// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

/*
Package x32 provides remote control support for the Behringer X32 Digital Mixer
using Open Sound Control remote protocol.
*/
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

// WriteMessage writes the OSC message.
func (m Mixer) WriteMessage(addr, typeTag string, args ...interface{}) error {
	msg, err := osc.Message(addr, typeTag, args...)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}

// MuteChannel mutes the given channel.
func (m Mixer) MuteChannel(ch int) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	addr := fmt.Sprintf("/ch/%02d/mix/on", ch)
	return m.WriteMessage(addr, "i", 0)
}

// MuteMain mutes the main channel.
func (m Mixer) MuteMain() error {
	return m.WriteMessage("/main/st/mix/on", "i", 0)
}

// NameChannel sets the name of the given channel. The name can only be up to
// 12 characters.
func (m Mixer) NameChannel(ch int, name string) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	if len(name) > 12 {
		return fmt.Errorf("channel name %s too long (12 char limit)", name)
	}
	addr := fmt.Sprintf("/ch/%02d/config/name", ch)
	return m.WriteMessage(addr, "s", name)
}

// UnmuteChannel unmutes the given channel.
func (m Mixer) UnmuteChannel(ch int) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	addr := fmt.Sprintf("/ch/%02d/mix/on", ch)
	return m.WriteMessage(addr, "i", 1)
}

// UnmuteMain unmutes the main channel.
func (m Mixer) UnmuteMain() error {
	return m.WriteMessage("/main/st/mix/on", "i", 1)
}

func validChannelRange(ch int) bool {
	return ch > 0 && ch < 33
}
