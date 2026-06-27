package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NumberOfAntennaPortsEUTRAAn1 int64 = 0
	NumberOfAntennaPortsEUTRAAn2 int64 = 1
	NumberOfAntennaPortsEUTRAAn4 int64 = 2
)

var numberOfAntennaPortsEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type NumberOfAntennaPortsEUTRA struct {
	Value int64
}

func (ie *NumberOfAntennaPortsEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, numberOfAntennaPortsEUTRAConstraints)
}

func (ie *NumberOfAntennaPortsEUTRA) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(numberOfAntennaPortsEUTRAConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
