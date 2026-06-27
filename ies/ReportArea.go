package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReportAreaCell int64 = 0
)

var reportAreaConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ReportArea struct {
	Value int64
}

func (ie *ReportArea) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, reportAreaConstraints)
}

func (ie *ReportArea) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(reportAreaConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
