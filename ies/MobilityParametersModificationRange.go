package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mobilityParametersModificationRangeConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "handoverTriggerChangeLowerLimit"},
		{Name: "handoverTriggerChangeUpperLimit"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MobilityParametersModificationRange struct {
	HandoverTriggerChangeLowerLimit int64
	HandoverTriggerChangeUpperLimit int64
	IEExtensions                    []byte
}

func (ie *MobilityParametersModificationRange) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityParametersModificationRangeConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.HandoverTriggerChangeLowerLimit, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(-20)),
		UpperBound: common.Ptr(int64(20)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.HandoverTriggerChangeUpperLimit, per.IntegerConstraints{
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

func (ie *MobilityParametersModificationRange) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityParametersModificationRangeConstraints)
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
		ie.HandoverTriggerChangeLowerLimit = val
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
		ie.HandoverTriggerChangeUpperLimit = val
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
