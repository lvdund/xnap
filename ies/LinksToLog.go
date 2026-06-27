package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LinksToLogUplink                int64 = 0
	LinksToLogDownlink              int64 = 1
	LinksToLogBothUplinkAndDownlink int64 = 2
)

var linksToLogConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type LinksToLog struct {
	Value int64
}

func (ie *LinksToLog) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, linksToLogConstraints)
}

func (ie *LinksToLog) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(linksToLogConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
