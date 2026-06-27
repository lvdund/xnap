package ies

import (
	"github.com/lvdund/asn1go/per"
)

var conditionalReconfigItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pCell-ID"},
		{Name: "pSCell-ID", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ConditionalReconfigItem struct {
	PCellID      TargetCGI
	PSCellID     *NRCGI
	IEExtensions []byte
}

func (ie *ConditionalReconfigItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(conditionalReconfigItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PSCellID != nil, false}); err != nil {
		return err
	}
	if err := ie.PCellID.Encode(e); err != nil {
		return err
	}
	if ie.PSCellID != nil {
		if err := ie.PSCellID.Encode(e); err != nil {
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

func (ie *ConditionalReconfigItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(conditionalReconfigItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PCellID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PSCellID = new(NRCGI)
		if err := ie.PSCellID.Decode(d); err != nil {
			return err
		}
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
