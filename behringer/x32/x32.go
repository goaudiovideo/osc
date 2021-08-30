// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package x32

import "io"

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
