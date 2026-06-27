package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ExcessPacketDelayThresholdValueMs0dot25 int64 = 0
	ExcessPacketDelayThresholdValueMs0dot5  int64 = 1
	ExcessPacketDelayThresholdValueMs1      int64 = 2
	ExcessPacketDelayThresholdValueMs2      int64 = 3
	ExcessPacketDelayThresholdValueMs4      int64 = 4
	ExcessPacketDelayThresholdValueMs5      int64 = 5
	ExcessPacketDelayThresholdValueMs10     int64 = 6
	ExcessPacketDelayThresholdValueMs20     int64 = 7
	ExcessPacketDelayThresholdValueMs30     int64 = 8
	ExcessPacketDelayThresholdValueMs40     int64 = 9
	ExcessPacketDelayThresholdValueMs50     int64 = 10
	ExcessPacketDelayThresholdValueMs60     int64 = 11
	ExcessPacketDelayThresholdValueMs70     int64 = 12
	ExcessPacketDelayThresholdValueMs80     int64 = 13
	ExcessPacketDelayThresholdValueMs90     int64 = 14
	ExcessPacketDelayThresholdValueMs100    int64 = 15
	ExcessPacketDelayThresholdValueMs150    int64 = 16
	ExcessPacketDelayThresholdValueMs300    int64 = 17
	ExcessPacketDelayThresholdValueMs500    int64 = 18
)

var excessPacketDelayThresholdValueConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
	ExtValues:  nil,
}

type ExcessPacketDelayThresholdValue struct {
	Value int64
}

func (ie *ExcessPacketDelayThresholdValue) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, excessPacketDelayThresholdValueConstraints)
}

func (ie *ExcessPacketDelayThresholdValue) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(excessPacketDelayThresholdValueConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
