package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	BandwidthMhz10  int64 = 0
	BandwidthMhz20  int64 = 1
	BandwidthMhz40  int64 = 2
	BandwidthMhz60  int64 = 3
	BandwidthMhz80  int64 = 4
	BandwidthMhz100 int64 = 5
)

var bandwidthConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  []int64{5},
}

type Bandwidth struct {
	Value int64
}

func (ie *Bandwidth) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, bandwidthConstraints)
}

func (ie *Bandwidth) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(bandwidthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
