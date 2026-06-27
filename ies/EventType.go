package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EventTypeReportUponChangeOfServingCell                      int64 = 0
	EventTypeReportUEMovingPresenceIntoOrOutOfTheAreaOfInterest int64 = 1
	EventTypeReportUponChangeOfServingCellAndAreaOfInterest     int64 = 2
)

var eventTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2},
}

type EventType struct {
	Value int64
}

func (ie *EventType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eventTypeConstraints)
}

func (ie *EventType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eventTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
