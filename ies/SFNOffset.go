package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sFNOffsetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sFN-Time-Offset"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SFNOffset struct {
	SFNTimeOffset per.BitString
	IEExtensions  []byte
}

func (ie *SFNOffset) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sFNOffsetConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.SFNTimeOffset, per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(24)),
		Max:        common.Ptr(int64(24)),
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

func (ie *SFNOffset) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sFNOffsetConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(24)),
			Max:        common.Ptr(int64(24)),
		})
		if err != nil {
			return err
		}
		ie.SFNTimeOffset = val
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
