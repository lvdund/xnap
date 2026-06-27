package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uESliceMaximumBitRateItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "s-NSSAI"},
		{Name: "dl-UE-Slice-MBR"},
		{Name: "ul-UE-Slice-MBR"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UESliceMaximumBitRateItem struct {
	SNSSAI       SNSSAI
	DlUESliceMBR BitRate
	UlUESliceMBR BitRate
	IEExtensions []byte
}

func (ie *UESliceMaximumBitRateItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uESliceMaximumBitRateItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SNSSAI.Encode(e); err != nil {
		return err
	}
	if err := ie.DlUESliceMBR.Encode(e); err != nil {
		return err
	}
	if err := ie.UlUESliceMBR.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UESliceMaximumBitRateItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uESliceMaximumBitRateItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SNSSAI.Decode(d); err != nil {
		return err
	}
	if err := ie.DlUESliceMBR.Decode(d); err != nil {
		return err
	}
	if err := ie.UlUESliceMBR.Decode(d); err != nil {
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
