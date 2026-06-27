package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MBSAssistanceInformationTrue int64 = 0
)

var mBSAssistanceInformationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MBSAssistanceInformation struct {
	Value int64
}

func (ie *MBSAssistanceInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mBSAssistanceInformationConstraints)
}

func (ie *MBSAssistanceInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mBSAssistanceInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
