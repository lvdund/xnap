package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var aSSecurityInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "key-NG-RAN-Star"},
		{Name: "ncc"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ASSecurityInformation struct {
	KeyNGRANStar per.BitString
	Ncc          int64
	IEExtensions []byte
}

func (ie *ASSecurityInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(aSSecurityInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.KeyNGRANStar, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(256)),
		Max:        common.Ptr(int64(256)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.Ncc, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(7)),
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

func (ie *ASSecurityInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(aSSecurityInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(256)),
			Max:        common.Ptr(int64(256)),
		})
		if err != nil {
			return err
		}
		ie.KeyNGRANStar = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(7)),
		})
		if err != nil {
			return err
		}
		ie.Ncc = val
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
