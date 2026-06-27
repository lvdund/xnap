package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RLCDuplicationInformationRLCPrimaryIndicatorTrue  int64 = 0
	RLCDuplicationInformationRLCPrimaryIndicatorFalse int64 = 1
)

var rLCDuplicationInformationRLCPrimaryIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RLCDuplicationInformationRLCPrimaryIndicator struct {
	Value int64
}

func (ie *RLCDuplicationInformationRLCPrimaryIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rLCDuplicationInformationRLCPrimaryIndicatorConstraints)
}

func (ie *RLCDuplicationInformationRLCPrimaryIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rLCDuplicationInformationRLCPrimaryIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var rLCDuplicationInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rLCDuplicationStateList"},
		{Name: "rLC-PrimaryIndicator", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type RLCDuplicationInformation struct {
	RLCDuplicationStateList RLCDuplicationStateList
	RLCPrimaryIndicator     *RLCDuplicationInformationRLCPrimaryIndicator
	IEExtensions            []byte
}

func (ie *RLCDuplicationInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCDuplicationInformationConstraints)
	if err := seq.EncodePreamble([]bool{ie.RLCPrimaryIndicator != nil, false}); err != nil {
		return err
	}
	if err := ie.RLCDuplicationStateList.Encode(e); err != nil {
		return err
	}
	if ie.RLCPrimaryIndicator != nil {
		if err := ie.RLCPrimaryIndicator.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RLCDuplicationInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCDuplicationInformationConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RLCDuplicationStateList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.RLCPrimaryIndicator = new(RLCDuplicationInformationRLCPrimaryIndicator)
		if err := ie.RLCPrimaryIndicator.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
