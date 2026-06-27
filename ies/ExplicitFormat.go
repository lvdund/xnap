package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var explicitFormatConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "permutation"},
		{Name: "noofDownlinkSymbols", Optional: true},
		{Name: "noofUplinkSymbols", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExplicitFormat struct {
	Permutation         Permutation
	NoofDownlinkSymbols *int64
	NoofUplinkSymbols   *int64
	IEExtensions        []byte
}

func (ie *ExplicitFormat) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(explicitFormatConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NoofDownlinkSymbols != nil, ie.NoofUplinkSymbols != nil, false}); err != nil {
		return err
	}
	if err := ie.Permutation.Encode(e); err != nil {
		return err
	}
	if ie.NoofDownlinkSymbols != nil {
		if err := e.EncodeInteger(*ie.NoofDownlinkSymbols, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(14)),
		}); err != nil {
			return err
		}
	}
	if ie.NoofUplinkSymbols != nil {
		if err := e.EncodeInteger(*ie.NoofUplinkSymbols, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(14)),
		}); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ExplicitFormat) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(explicitFormatConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Permutation.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(14)),
		})
		if err != nil {
			return err
		}
		ie.NoofDownlinkSymbols = &val
	}
	if seq.IsComponentPresent(2) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(14)),
		})
		if err != nil {
			return err
		}
		ie.NoofUplinkSymbols = &val
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
