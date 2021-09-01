// Copyright (c) 2021 The goaudiovideo developers. All rights reserved.
// Project site: https://github.com/goaudiovideo/osc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package x32

// Icon provudes an enumeration for the avialable icons.
type Icon int

// Enum for X32 icons.
const (
	EmptyIcon            Icon = 1
	BassKickDrum              = 2
	BassDrum                  = 3
	SticksAboveSnareDrum      = 4
	SnareDrumAboveSticks      = 5
	HighTomDrum               = 6
	MediumTomDrum             = 7
	FloorTomDrum              = 8
	HighHatCymbal             = 9
	Cymbal                    = 10
	DrumSet                   = 11
	Lights                    = 12
	Bongos                    = 13
	Bongos2                   = 14
	Tambourine                = 15
	Xylophone                 = 16
	Guitar1                   = 17
	Guitar2                   = 18
	Guitar3                   = 19
	Guitar4                   = 20
	Guitar5                   = 21
	Guitar6                   = 22
	Guitar7                   = 23
	Speaker1                  = 24
	Speaker2                  = 25
	Speaker3                  = 26
	GrandPiano                = 27
	UprightPiano              = 28
	Keyboard1                 = 29
	Keyboard2                 = 30
	Keyboard3                 = 31
	Keyboard4                 = 32
	Keyboard5                 = 33
	Keyboard6                 = 34
	Trumpet                   = 35
	Trombone                  = 36
	Saxaphone                 = 37
	Clarinet                  = 38
	Violin                    = 39
	Cello                     = 40
	MaleSinger                = 41
	FemaleSinger              = 42
	Singers                   = 43
	Rockon                    = 44
	ATalker                   = 45
	BTalker                   = 46
	InEarMic                  = 53
	Laptop                    = 62
	SmileyFace                = 74
)

// String implements the Stringer interface for Icon.
func (icon Icon) String() string {
	return iconDescription[icon]
}

var iconDescription = map[Icon]string{
	EmptyIcon:            "empty",
	BassKickDrum:         "bass kick drum",
	BassDrum:             "bass drum",
	SticksAboveSnareDrum: "sticks above snare drum",
	SnareDrumAboveSticks: "snare drum above sticks",
	HighTomDrum:          "high tom drum",
	MediumTomDrum:        "medium tom drum",
	FloorTomDrum:         "floor tom drum",
	HighHatCymbal:        "high hat cymbals",
	Cymbal:               "cymbals",
	DrumSet:              "full drum set",
	Lights:               "lights",
	Bongos:               "bongos",
	Bongos2:              "bongos 2",
	Tambourine:           "tambourine",
	Xylophone:            "xylophone",
	Guitar1:              "guitar 1",
	Guitar2:              "guitar 2",
	Guitar3:              "guitar 3",
	Guitar4:              "guitar 4",
	Guitar5:              "guitar 5",
	Guitar6:              "guitar 6",
	Guitar7:              "guitar 7",
	Speaker1:             "speakers 1",
	Speaker2:             "speakers 2",
	Speaker3:             "speakers 3",
	GrandPiano:           "grand piano",
	UprightPiano:         "upright piano",
	Keyboard1:            "keyboard 1",
	Keyboard2:            "keyboard 2",
	Keyboard3:            "keyboard 3",
	Keyboard4:            "keyboard 4",
	Keyboard5:            "keyboard 5",
	Keyboard6:            "keyboard 6",
	Trumpet:              "trumpet",
	Trombone:             "trombone",
	Saxaphone:            "saxaphone",
	Clarinet:             "clarinet",
	Violin:               "violin",
	Cello:                "cello",
	MaleSinger:           "male singer",
	FemaleSinger:         "female singer",
	Singers:              "trio of singers",
	Rockon:               "rock on hand gesture",
	ATalker:              "A talker",
	BTalker:              "B talker",
	InEarMic:             "in-ear microphone",
	Laptop:               "laptop",
	SmileyFace:           "smiley face",
}
