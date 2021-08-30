// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package osc

// Message creates an OSC message and returns it as a byte slice.
func Message(addr, typeTag string, args ...interface{}) ([]byte, error) {
	// Add OSC Address Pattern to message.
	msg := []byte(addr)
	msg = append(msg, 0)
	msg = addPadBytes(msg)

	// Add OSC Type Tag to message or add a comma and nulls if there aren't any
	// type tags provided.
	msg = append(msg, []byte(",")...)
	if typeTag != "" {
		msg = append(msg, typeTag...)
	}
	msg = addPadBytes(msg)
	return msg, nil
}

// addPadBytes adds the proper number of pad bytes.
func addPadBytes(msg []byte) []byte {
	n := numPadBytes(len(msg))
	for i := 0; i < n; i++ {
		msg = append(msg, 0)
	}
	return msg
}

// numPadBytes calculates the number of pad bytes required to make a message
// have a byte count that's a multiple of four.
func numPadBytes(l int) int {
	return (4 - (l % 4)) % 4
}
