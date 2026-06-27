package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBsToBeActivatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbIndex"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SSBsToBeActivatedItem struct {
	SsbIndex     int64
	IEExtensions []byte
}

func (ie *SSBsToBeActivatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBsToBeActivatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SsbIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(63)),
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

func (ie *SSBsToBeActivatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBsToBeActivatedItemConstraints)
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
		ie.SsbIndex = val
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
