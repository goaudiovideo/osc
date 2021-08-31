// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package osc

import (
	"bytes"
	"encoding/binary"
)

// Message creates an OSC message and returns it as a byte slice.
func Message(addr, typeTag string, args ...interface{}) ([]byte, error) {
	// Add OSC Address Pattern to message and appropriate number of zero bytes.
	msg := []byte(addr)
	msg = append(msg, 0)
	msg = addZeroBytes(msg)

	// Add OSC Type Tag to message or add a comma and nulls if there aren't any
	// type tags provided and the appropriate number of zero bytes.
	msg = append(msg, []byte(",")...)
	if typeTag != "" {
		msg = append(msg, typeTag...)
	}
	msg = addZeroBytes(msg)

	// Add args to message if there are any given.
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			msg = append(msg, arg...)
			msg = append(msg, 0)
		case float32:
			b, err := encodeFloat32(arg)
			if err != nil {
				return nil, err
			}
			msg = append(msg, b...)
		}
		msg = addZeroBytes(msg)
	}

	return msg, nil
}

// addZeroBytes adds the proper number of zero bytes.
func addZeroBytes(msg []byte) []byte {
	n := numZeroBytes(len(msg))
	if n > 0 {
		padBytes := make([]byte, n)
		msg = append(msg, padBytes...)
	}
	return msg
}

// numZeroBytes calculates the number of zeo bytes to append to a message in
// order to have a byte count that's a multiple of four.
func numZeroBytes(l int) int {
	return (4 - (l % 4)) % 4
}

// encodeFloat32 converts a float32 number into the big-endian binary byte
// slice required by an OSC message.
func encodeFloat32(f float32) ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
