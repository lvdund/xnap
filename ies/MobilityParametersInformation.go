package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mobilityParametersInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "handoverTriggerChange"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MobilityParametersInformation struct {
	HandoverTriggerChange int64
	IEExtensions          []byte
}

func (ie *MobilityParametersInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityParametersInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.HandoverTriggerChange, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(-20)),
		UpperBound: common.Ptr(int64(20)),
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MobilityParametersInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityParametersInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(-20)),
			UpperBound: common.Ptr(int64(20)),
		})
		if err != nil {
			return err
		}
		ie.HandoverTriggerChange = val
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
