package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PSCellChangeHistoryReportingFullHistory int64 = 0
)

var pSCellChangeHistoryConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PSCellChangeHistory struct {
	Value int64
}

func (ie *PSCellChangeHistory) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pSCellChangeHistoryConstraints)
}

func (ie *PSCellChangeHistory) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pSCellChangeHistoryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
