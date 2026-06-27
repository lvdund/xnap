package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var symbolAllocationInSlotBothDLandULConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "numberofDLSymbols"},
		{Name: "numberofULSymbols"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SymbolAllocationInSlotBothDLandUL struct {
	NumberofDLSymbols int64
	NumberofULSymbols int64
	IEExtensions      []byte
}

func (ie *SymbolAllocationInSlotBothDLandUL) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(symbolAllocationInSlotBothDLandULConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.NumberofDLSymbols, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(13)),
	}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.NumberofULSymbols, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(13)),
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

func (ie *SymbolAllocationInSlotBothDLandUL) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(symbolAllocationInSlotBothDLandULConstraints)
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
			UpperBound: common.Ptr(int64(13)),
		})
		if err != nil {
			return err
		}
		ie.NumberofDLSymbols = val
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(13)),
		})
		if err != nil {
			return err
		}
		ie.NumberofULSymbols = val
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
