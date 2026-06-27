package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var slotConfigurationListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotIndex"},
		{Name: "symbolAllocation-in-Slot"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SlotConfigurationListItem struct {
	SlotIndex              int64
	SymbolAllocationInSlot SymbolAllocationInSlot
	IEExtensions           []byte
}

func (ie *SlotConfigurationListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotConfigurationListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SlotIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(5119)),
	}); err != nil {
		return err
	}
	if err := ie.SymbolAllocationInSlot.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SlotConfigurationListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotConfigurationListItemConstraints)
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
			UpperBound: common.Ptr(int64(5119)),
		})
		if err != nil {
			return err
		}
		ie.SlotIndex = val
	}
	if err := ie.SymbolAllocationInSlot.Decode(d); err != nil {
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
