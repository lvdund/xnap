package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EventTypeTriggerOutOfCoverageTrue int64 = 0
)

var eventTypeTriggerOutOfCoverageConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type EventTypeTriggerOutOfCoverage struct {
	Value int64
}

func (ie *EventTypeTriggerOutOfCoverage) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eventTypeTriggerOutOfCoverageConstraints)
}

func (ie *EventTypeTriggerOutOfCoverage) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eventTypeTriggerOutOfCoverageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	EventTypeTriggerChOutOfCoverage    = 0
	EventTypeTriggerChEventL1          = 1
	EventTypeTriggerChChoiceExtensions = 2
)

var eventTypeTriggerConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "outOfCoverage"},
		{Name: "eventL1"},
		{Name: "choice-Extensions"},
	},
	ExtAlternatives: nil,
}

type EventTypeTrigger struct {
	Choice          int
	OutOfCoverage   *EventTypeTriggerOutOfCoverage
	EventL1         *EventL1
	ChoiceExtension []byte
}

func (ie *EventTypeTrigger) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(eventTypeTriggerConstraints)
	switch ie.Choice {
	case 0: // outOfCoverage
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.OutOfCoverage.Encode(e); err != nil {
			return err
		}
	case 1: // eventL1
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.EventL1.Encode(e); err != nil {
			return err
		}
	case 2: // choice-Extensions
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtensions (kind=ext)
	}
	return nil
}

func (ie *EventTypeTrigger) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(eventTypeTriggerConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // outOfCoverage
		ie.OutOfCoverage = new(EventTypeTriggerOutOfCoverage)
		if err := ie.OutOfCoverage.Decode(d); err != nil {
			return err
		}
	case 1: // eventL1
		ie.EventL1 = new(EventL1)
		if err := ie.EventL1.Decode(d); err != nil {
			return err
		}
	case 2: // choice-Extensions
		// TODO decode field ChoiceExtensions (kind=ext)
	}
	return nil
}
