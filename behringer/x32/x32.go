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

// MuteChannel mutes the given channel.
func (m Mixer) MuteChannel(ch int) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	addr := fmt.Sprintf("/ch/%02d/mix/on", ch)
	msg, err := osc.Message(addr, "i", 0)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}

// MuteMain mutes the main channel.
func (m Mixer) MuteMain() error {
	msg, err := osc.Message("/main/st/mix/on", "i", 0)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
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
	msg, err := osc.Message(addr, "s", name)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}

// UnmuteChannel unmutes the given channel.
func (m Mixer) UnmuteChannel(ch int) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	addr := fmt.Sprintf("/ch/%02d/mix/on", ch)
	msg, err := osc.Message(addr, "i", 1)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}

// UnmuteMain unmutes the main channel.
func (m Mixer) UnmuteMain() error {
	msg, err := osc.Message("/main/st/mix/on", "i", 1)
	if err != nil {
		return err
	}
	_, err = m.Write(msg)
	return err
}

func validChannelRange(ch int) bool {
	return ch > 0 && ch < 33
}
