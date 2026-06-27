package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var n6JitterInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "n6JitterLowerBound"},
		{Name: "n6JitterUpperBound"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type N6JitterInformation struct {
	N6JitterLowerBound int64
	N6JitterUpperBound int64
	IEExtensions       []byte
}

func (ie *N6JitterInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(n6JitterInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.N6JitterLowerBound, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(-127)),
		UpperBound: common.Ptr(int64(127)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.N6JitterUpperBound, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(-127)),
		UpperBound: common.Ptr(int64(127)),
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

func (ie *N6JitterInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(n6JitterInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(-127)),
			UpperBound: common.Ptr(int64(127)),
		})
		if err != nil {
			return err
		}
		ie.N6JitterLowerBound = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(-127)),
			UpperBound: common.Ptr(int64(127)),
		})
		if err != nil {
			return err
		}
		ie.N6JitterUpperBound = val
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
