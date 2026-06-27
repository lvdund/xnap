package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCPACInterSNExecutionNotifyExecuted int64 = 0
)

var sCPACInterSNExecutionNotifyConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SCPACInterSNExecutionNotify struct {
	Value int64
}

func (ie *SCPACInterSNExecutionNotify) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCPACInterSNExecutionNotifyConstraints)
}

func (ie *SCPACInterSNExecutionNotify) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCPACInterSNExecutionNotifyConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
