package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LoggingDurationM10  int64 = 0
	LoggingDurationM20  int64 = 1
	LoggingDurationM40  int64 = 2
	LoggingDurationM60  int64 = 3
	LoggingDurationM90  int64 = 4
	LoggingDurationM120 int64 = 5
)

var loggingDurationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2, 3, 4, 5},
	ExtValues:  nil,
}

type LoggingDuration struct {
	Value int64
}

func (ie *LoggingDuration) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, loggingDurationConstraints)
}

func (ie *LoggingDuration) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(loggingDurationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
