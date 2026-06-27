package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FrequencyShift7p5khzFalse int64 = 0
	FrequencyShift7p5khzTrue  int64 = 1
)

var frequencyShift7p5khzConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FrequencyShift7p5khz struct {
	Value int64
}

func (ie *FrequencyShift7p5khz) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, frequencyShift7p5khzConstraints)
}

func (ie *FrequencyShift7p5khz) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(frequencyShift7p5khzConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
