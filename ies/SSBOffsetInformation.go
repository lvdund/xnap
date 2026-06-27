package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBOffsetInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sSBIndex"},
		{Name: "sSBTriggeringOffset"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SSBOffsetInformation struct {
	SSBIndex            int64
	SSBTriggeringOffset MobilityParametersInformation
	IEExtensions        []byte
}

func (ie *SSBOffsetInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBOffsetInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SSBIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(63)),
	}); err != nil {
		return err
	}
	if err := ie.SSBTriggeringOffset.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SSBOffsetInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBOffsetInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(63)),
		})
		if err != nil {
			return err
		}
		ie.SSBIndex = val
	}
	if err := ie.SSBTriggeringOffset.Decode(d); err != nil {
		return err
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
