// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

/*
Package x32 provides remote control support for the Behringer X32 Digital Mixer
using Open Sound Control (OSC) remote protocol.
*/
package x32

import (
	"bufio"
	"fmt"
	"io"

	"github.com/goaudiovideo/osc"
)

// Info models the info received back from the mixer.
type Info struct {
	ServerVersion  string
	ServerName     string
	ConsoleModel   string
	ConsoleVersion string
}

// Mixer models a Behringer X32 mixer that can be controlled using Open Sound
// Control (OSC).
type Mixer struct {
	rw io.ReadWriter
}

// NewMixer creates a new Mixer using the given writer.
func NewMixer(rw io.ReadWriter) Mixer {
	return Mixer{
		rw: rw,
	}
}

// Write implements the Writer interface for Mixer.
func (m Mixer) Write(p []byte) (int, error) {
	return m.rw.Write(p)
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

// Info returns information about the X32 Mixer.
// FIXME(mdr): Not working! Need to create functions to handle reading OSC
// messages.
func (m Mixer) Info() (Info, error) {
	info := Info{}
	err := m.WriteMessage("/info", "s")
	if err != nil {
		return info, err
	}
	p := make([]byte, 128)
	_, err = bufio.NewReader(m.rw).Read(p)
	if err != nil {
		return info, err
	}
	return info, nil
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

// SetChannelColor sets the color for the given channel.
func (m Mixer) SetChannelColor(ch int, color Color) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	addr := fmt.Sprintf("/ch/%02d/config/color", ch)
	return m.WriteMessage(addr, "i", int(color))
}

// SetChannelIcon sets the icon for the given channel.
func (m Mixer) SetChannelIcon(ch int, icon Icon) error {
	if !validChannelRange(ch) {
		return fmt.Errorf("channel %d out of range 1-32", ch)
	}
	addr := fmt.Sprintf("/ch/%02d/config/icon", ch)
	return m.WriteMessage(addr, "i", int(icon))
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

// dbLevelToDecimal converts a dB level from -90.0 dB to +10.0 dB to a decimal
// value in the range of 0.0 to 1.0.
func dbLevelToDecimal(db float64) float64 {
	switch {
	case db < -90.0:
		return 0.0
	case db < -60.0:
		return (db + 90.0) / 480.0
	case db < -30.0:
		return (db + 70.0) / 160.0
	case db < -10.0:
		return (db + 50.0) / 80.0
	case db <= 10.0:
		return (db + 30.0) / 40.0
	default:
		return 1.0
	}
}

// decimalToDBLevel converts a decimal value in the range of 0.0 to 1.0 to a dB
// level from -90.0 dB to +10.0.
func decimalToDBLevel(f float64) float64 {
	switch {
	case f > 1.0:
		return 10.0
	case f >= 0.5:
		return f*40.0 - 30.0
	case f >= 0.25:
		return f*80.0 - 50.0
	case f >= 0.0625:
		return f*160.0 - 70.0
	case f >= 0.0:
		return f*480.0 - 90.
	default:
		return -90.0
	}
}
